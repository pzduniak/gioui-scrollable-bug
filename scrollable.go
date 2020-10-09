package main

import (
	"image"

	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"git.sr.ht/~whereswaldon/scroll"
)

const inf = 1e6

type Scrollable struct {
	list      *layout.List
	indicator *scroll.Scrollable
	gesture   gesture.Scroll

	currentOffset      int
	currentHeight      int
	currentTotalHeight int
}

func NewScrollable(
	alignment layout.Alignment,
	scrollToEnd bool,
) *Scrollable {
	return &Scrollable{
		indicator: &scroll.Scrollable{},
	}
}

func (s *Scrollable) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	/*
		list := &layout.List{
			Axis: layout.Vertical,
		}
		return list.Layout(gtx, 1, func(gtx layout.Context, idx int) layout.Dimensions {
			return widget(gtx)
		})
	*/

	maxHeight := gtx.Constraints.Max.Y
	renderingScrollbar := s.currentHeight < s.currentTotalHeight

	scrollbarPadding := unit.Dp(0)
	scrollbarInnerPadding := unit.Dp(0)
	if renderingScrollbar {
		scrollbarPadding = unit.Dp(2)
		scrollbarInnerPadding = unit.Dp(2)
	}

	flexChildren := []layout.FlexChild{
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: scrollbarInnerPadding,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// this component always scrolls vertically, so we only overwrite the Y limits
				gtx.Constraints = layout.Constraints{
					Min: image.Pt(gtx.Constraints.Min.X, 0),
					Max: image.Pt(gtx.Constraints.Max.X, inf),
				}
				stack := op.Push(gtx.Ops)
				macro := op.Record(gtx.Ops)
				dims := widget(gtx)
				call := macro.Stop()
				stack.Pop()

				s.currentTotalHeight = dims.Size.Y

				// scrollbar logic
				if d := s.gesture.Scroll(gtx.Metric, gtx, gtx.Now, gesture.Vertical); d != 0 {
					s.currentOffset += d
				}
				if scrolled, progress := s.indicator.Scrolled(); scrolled {
					s.currentOffset = int(float32(dims.Size.Y) * progress)
				}
				// gesture logic
				if s.currentOffset+s.currentHeight > s.currentTotalHeight {
					s.currentOffset = s.currentTotalHeight - s.currentHeight
					if s.gesture.State() == gesture.StateFlinging {
						s.gesture.Stop()
					}
				}
				if s.currentOffset < 0 {
					s.currentOffset = 0
					if s.gesture.State() == gesture.StateFlinging {
						s.gesture.Stop()
					}
				}

				op.Offset(f32.Pt(0, -float32(s.currentOffset))).Add(gtx.Ops)
				clip.Rect(image.Rectangle{
					Min: image.Pt(0, s.currentOffset),
					Max: image.Pt(gtx.Constraints.Max.X, maxHeight+s.currentOffset),
				}).Add(gtx.Ops)
				call.Add(gtx.Ops)

				dimsY := dims.Size.Y
				if dimsY > maxHeight {
					dimsY = maxHeight
				}
				s.currentHeight = dimsY

				return layout.Dimensions{
					Size: image.Pt(dims.Size.X, dimsY),
				}
			})
		}),
	}

	if renderingScrollbar {
		flexChildren = append(flexChildren,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return scroll.DefaultBar(
					s.indicator,
					float32(s.currentOffset)/float32(s.currentTotalHeight), // depth in items count we scrolled past/see
					float32(s.currentHeight)/float32(s.currentTotalHeight), // fraction of the visible items visible on the screen
				).Layout(gtx)
			}),
		)
	}

	dims := layout.Inset{
		Right: scrollbarPadding,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis:      layout.Horizontal,
			Spacing:   layout.SpaceBetween,
			Alignment: layout.Start,
		}.Layout(
			gtx,
			flexChildren...,
		)
	})
	pointer.Rect(image.Rect(0, 0, dims.Size.X, dims.Size.Y)).Add(gtx.Ops)
	s.gesture.Add(gtx.Ops)
	return dims
}

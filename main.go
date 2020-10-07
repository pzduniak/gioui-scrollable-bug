package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/dchest/uniuri"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

var (
	scroll1 = NewScrollable(layout.Start, false)
	scroll2 = NewScrollable(layout.Start, false)
	scroll3 = NewScrollable(layout.Start, false)
	scroll4 = NewScrollable(layout.Start, false)

	copyField1 = NewCopyField()
	copyField2 = NewCopyField()
	copyField3 = NewCopyField()
	copyField4 = NewCopyField()
	copyField5 = NewCopyField()

	copyVals = []string{
		uniuri.New(),
		uniuri.New(),
		uniuri.New(),
		uniuri.New(),
		uniuri.New(),
	}
)

func main() {
	go func() {
		window := app.NewWindow(
			app.Size(unit.Dp(600), unit.Dp(300)),
			app.Title("Bug demo"),
		)
		th := material.NewTheme(gofont.Collection())
		var ops op.Ops
		for {
			select {
			case e := <-window.Events():
				switch e := e.(type) {
				case system.DestroyEvent:
					panic(e.Err)
				case system.FrameEvent:
					gtx := layout.NewContext(&ops, e)
					frame(gtx, th)
					e.Frame(gtx.Ops)
				}
			}
		}
	}()
	app.Main()
}

func frame(gtx C, th *material.Theme) D {
	return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
		return layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(
			gtx,
			layout.Flexed(1, func(gtx C) D {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(
					gtx,
					layout.Flexed(1, func(gtx C) D {
						return layout.Inset{
							Right: unit.Dp(4),
						}.Layout(gtx, func(gtx C) D {
							return frameSection(gtx, th, scroll1)
						})
					}),
					layout.Flexed(1, func(gtx C) D {
						return layout.Inset{
							Right: unit.Dp(4),
						}.Layout(gtx, func(gtx C) D {
							return frameSection(gtx, th, scroll2)
						})
					}),
				)
			}),
			layout.Flexed(1, func(gtx C) D {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(
					gtx,
					layout.Flexed(1, func(gtx C) D {
						return layout.Inset{
							Right: unit.Dp(4),
						}.Layout(gtx, func(gtx C) D {
							return frameSection(gtx, th, scroll3)
						})
					}),
					layout.Flexed(1, func(gtx C) D {
						return layout.Inset{
							Right: unit.Dp(4),
						}.Layout(gtx, func(gtx C) D {
							return frameSection(gtx, th, scroll4)
						})
					}),
				)
			}),
		)
	})
}

func frameSection(gtx C, th *material.Theme, sc *Scrollable) D {
	rowHeight := gtx.Px(unit.Dp(40))
	wc := WestCenter{
		Height: rowHeight,
	}

	return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
		return sc.Layout(gtx, func(gtx C) D {
			return layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(
				gtx,
				layout.Rigid(func(gtx C) D {
					// left side is all labels
					return layout.Inset{
						Right: unit.Dp(8),
					}.Layout(gtx, func(gtx C) D {
						return layout.Flex{
							Axis:      layout.Vertical,
							Spacing:   layout.SpaceEnd,
							Alignment: layout.Start,
						}.Layout(
							gtx,
							layout.Rigid(func(gtx C) D {
								return wc.Layout(gtx, material.Body1(th, "First:").Layout)
							}),
							layout.Rigid(func(gtx C) D {
								return wc.Layout(gtx, material.Body1(th, "Second:").Layout)
							}),
							layout.Rigid(func(gtx C) D {
								return wc.Layout(gtx, material.Body1(th, "Third:").Layout)
							}),
							layout.Rigid(func(gtx C) D {
								return wc.Layout(gtx, material.Body1(th, "Fourth:").Layout)
							}),
							layout.Rigid(func(gtx C) D {
								return wc.Layout(gtx, material.Body1(th, "Fifth:").Layout)
							}),
						)
					})
				}),
				layout.Flexed(1, func(gtx C) D {
					// right side is all copyfields
					return layout.Flex{
						Axis:      layout.Vertical,
						Spacing:   layout.SpaceEnd,
						Alignment: layout.Start,
					}.Layout(
						gtx,
						layout.Rigid(func(gtx C) D {
							return wc.Layout(gtx, func(gtx C) D {
								return copyField1.Layout(gtx, th, copyVals[0])
							})
						}),
						layout.Rigid(func(gtx C) D {
							return wc.Layout(gtx, func(gtx C) D {
								return copyField2.Layout(gtx, th, copyVals[1])
							})
						}),
						layout.Rigid(func(gtx C) D {
							return wc.Layout(gtx, func(gtx C) D {
								return copyField3.Layout(gtx, th, copyVals[2])
							})
						}),
						layout.Rigid(func(gtx C) D {
							return wc.Layout(gtx, func(gtx C) D {
								return copyField4.Layout(gtx, th, copyVals[3])
							})
						}),
						layout.Rigid(func(gtx C) D {
							return wc.Layout(gtx, func(gtx C) D {
								return copyField5.Layout(gtx, th, copyVals[4])
							})
						}),
					)
				}),
			)
		})
	})
}

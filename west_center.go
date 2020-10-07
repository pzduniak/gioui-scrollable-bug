package main

import (
	"gioui.org/layout"
)

type WestCenter struct {
	Height int
}

func (w WestCenter) Layout(
	gtx layout.Context,
	widget layout.Widget,
) layout.Dimensions {
	gtx.Constraints.Min.Y = w.Height
	gtx.Constraints.Max.Y = w.Height
	return layout.Flex{
		Axis:      layout.Horizontal,
		Spacing:   layout.SpaceEnd,
		Alignment: layout.Middle,
	}.Layout(
		gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.W.Layout(
				gtx,
				widget,
			)
		}),
	)
}

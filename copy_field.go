package main

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type CopyField struct {
	Button *widget.Clickable
	Editor *widget.Editor
}

func NewCopyField() *CopyField {
	return &CopyField{
		Button: &widget.Clickable{},
		Editor: &widget.Editor{SingleLine: true},
	}
}

func (w *CopyField) Process(window *app.Window) {
	for w.Button.Clicked() {
		window.WriteClipboard(w.Editor.Text())
	}
}

func (w *CopyField) Layout(gtx layout.Context, th *material.Theme, text string) layout.Dimensions {
	if w.Editor.Text() != text {
		w.Editor.SetText(text)
	}

	return layout.Flex{
		Axis:      layout.Horizontal,
		Spacing:   layout.SpaceBetween,
		Alignment: layout.Middle,
	}.Layout(
		gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max.Y = gtx.Px(unit.Dp(37))
			return layout.Inset{
				Right: unit.Dp(2),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return widget.Border{
					Color:        color.RGBA{A: 64},
					CornerRadius: unit.Dp(4),
					Width:        unit.Dp(1),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						editor := material.Editor(th, w.Editor, "")
						// there's currently no way to render it all with 100% opacity
						return editor.Layout(gtx.Disabled())
					})
				})
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left: unit.Dp(2),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				button := material.Button(th, w.Button, "Copy")
				button.Background = color.RGBA{14, 113, 235, 255}
				button.Color = color.RGBA{255, 255, 255, 255}
				return button.Layout(gtx)
			})
		}),
	)
}

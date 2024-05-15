package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

func main() {
	th := NewTheme(gofont.Collection())

	ui := UI{
		Window: new(app.Window),
		Theme:  th,
		Menu:   &Menu{},
		Resize: component.Resize{Ratio: 0.5},
	}

	go func() {
		if err := ui.Loop(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type (
	C = layout.Context
	D = layout.Dimensions
)

type UI struct {
	Window *app.Window
	Theme  *Theme
	Menu   *Menu
	component.Resize
}

type Theme struct {
	Base *material.Theme
}

func NewTheme(font []font.FontFace) *Theme {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(font))
	return &Theme{
		Base: th,
	}
}

func (ui UI) Loop() error {
	var ops op.Ops
	for {
		e := ui.Window.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			ui.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

func Overflow() []component.OverflowAction {
	return []component.OverflowAction{
		{
			Name: "Close",
		},
	}
}

func (ui *UI) Layout(gtx C) D {
	tiled := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return ui.Resize.Layout(gtx,
			func(gtx C) D {
				return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx C) D {
					c := color.NRGBA{R: 40, G: 40, B: 40, A: 255}
					paint.FillShape(gtx.Ops, c, clip.Rect{Max: gtx.Constraints.Max}.Op())
					return D{Size: gtx.Constraints.Max}

				})
			},
			func(gtx C) D {
				tile := layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx C) D {
					c := color.NRGBA{R: 40, G: 40, B: 40, A: 255}
					paint.FillShape(gtx.Ops, c, clip.Rect{Max: gtx.Constraints.Max}.Op())
					return D{Size: gtx.Constraints.Max}
				})
				return tile
			},
			func(gtx C) D {
				rect := image.Rectangle{
					Max: image.Point{
						X: (gtx.Dp(unit.Dp(3))),
						Y: (gtx.Constraints.Max.Y),
					},
				}
				paint.FillShape(gtx.Ops, color.NRGBA{A: 255}, clip.Rect(rect).Op())
				return D{Size: rect.Max}
			},
		)
	})
	barFlex := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		bar := ui.Menu.Layout(gtx)
		return bar
	})

	flex := layout.Flex{Axis: layout.Vertical}
	flex.Layout(gtx, barFlex, tiled)
	return layout.Dimensions{Size: gtx.Constraints.Max}
}

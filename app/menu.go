package main

import (
	"image"
	"image/color"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type MenuItem struct {
	Label    string
	Shortcut string
	Disabled bool
	Handler  func()
}

type Menu struct {
	Items   []MenuItem
	open    bool
	lastIdx int
}

func (m *Menu) Open() {
	m.open = true
}

func (m *Menu) Close() {
	m.open = false
	m.lastIdx = -1
}

func (m *Menu) Layout(gtx layout.Context) layout.Dimensions {
	if !m.open {
		return layout.Dimensions{}
	}

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
						layout.Flexed(0.3, m.menuIcon(gtx)),
						layout.Flexed(0.7, m.menuLabel(gtx)),
					)
				},
			)
		}),
	)
}

func (m *Menu) menuIcon(gtx layout.Context) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Top: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					paint.ColorOp{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 255}}.Add(gtx.Ops)
					paint.PaintOp{}.Add(gtx.Ops)
					return layout.Dimensions{Size: image.Point{X: 32, Y: 32}}
				})
			})
		})
	}
}

func (m *Menu) menuLabel(gtx layout.Context) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		label := material.Body1(nil, m.Items[m.lastIdx].Label)
		label.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		return layout.Center.Layout(gtx, label.Layout)
	}
}

func (m *Menu) Event(gtx layout.Context, ev interface{}) {
	if m.open {
		switch e := ev.(type) {
		case pointer.Event:
			// Handle pointer event here
			_ = e
		}
	}
}


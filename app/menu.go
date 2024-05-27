package main

import (
	"image/color"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/widget"
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
	Open    bool
	lastIdx int
}

func (m *Menu) Toggle() {
	m.Open = !m.Open
}

func (m *Menu) Layout(gtx layout.Context) layout.Dimensions {
	if !m.Open {
		return layout.Dimensions{}
	}

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical, Alignment: layout.Start}.Layout(gtx,
						m.menuItems(gtx)...,
					)
				},
			)
		}),
	)
}

func (m *Menu) menuItems(gtx layout.Context) []layout.FlexChild {
	var items []layout.FlexChild
	for i := range m.Items {
		i := i // capture loop variable
		item := m.Items[i]
		items = append(items, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Bottom: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(material.NewTheme(), new(widget.Clickable), item.Label)
				if item.Disabled {
					btn.Background = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
				}
				return btn.Layout(gtx)
			})
		}))
	}
	return items
}

func (m *Menu) Event(gtx layout.Context, ev interface{}) {
	if m.Open {
		switch e := ev.(type) {
		case pointer.Event:
			// Handle pointer event here
			_ = e
		}
	}
}

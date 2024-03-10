
package main

import (
	"fmt"
	"math"
)

type Container struct {
	WidthPercent, HeightPercent float64
	Children                    []*Container
}

type TiledView struct {
	Root *Container
}

func NewContainer(widthPercent, heightPercent float64) *Container {
	return &Container{
		WidthPercent:  widthPercent,
		HeightPercent: heightPercent,
	}
}

func NewTiledView(root *Container) *TiledView {
	return &TiledView{Root: root}
}

func (c *Container) AddChild(child *Container) {
	c.Children = append(c.Children, child)
}

func (tv *TiledView) TileInFibonacciLayout(container *Container, availableWidth, availableHeight float64) {
	if container == nil {
		return
	}

	numChildren := len(container.Children)
	if numChildren == 0 {
		return
	}

	fib := []int{1, 1}
	for fib[len(fib)-1] < numChildren {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
	}

	totalWidth := 0.0
	totalHeight := 0.0
	for _, child := range container.Children {
		totalWidth += child.WidthPercent
		totalHeight += child.HeightPercent
	}

	x := 0.0
	y := 0.0
	for i, child := range container.Children {
		child.WidthPercent = (child.WidthPercent / totalWidth) * availableWidth
		child.HeightPercent = (child.HeightPercent / totalHeight) * availableHeight
		fmt.Printf("Container %d: X=%.2f%%, Y=%.2f%%, Width=%.2f%%, Height=%.2f%%\n", i+1, x*100/availableWidth, y*100/availableHeight, child.WidthPercent*100/availableWidth, child.HeightPercent*100/availableHeight)
		if i < len(fib) {
			if i%2 == 0 {
				x += child.WidthPercent
			} else {
				y += child.HeightPercent
			}
		}
		tv.TileInFibonacciLayout(child, child.WidthPercent, child.HeightPercent)
	}
}


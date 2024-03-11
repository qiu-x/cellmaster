package layouts

import (
	"cellmaster/gui"
	"slices"
)

type GridLayout struct {
	parent        gui.INode
	grid          [][]gui.INode
	RowSizes      []float64
	ColumnSizes   []float64
	rowOffsets    []float64
	columnOffsets []float64
}

func (gl *GridLayout) Children() []gui.INode {
	return slices.Concat(gl.grid...)
}

func (gl *GridLayout) Parent() gui.INode {
	return gl.parent
}

func NewGridLayout(rows, columns int) *GridLayout {
	grid := make([][]gui.INode, rows)
	for i := range grid {
		grid[i] = make([]gui.INode, columns)
	}
	gridLayout := &GridLayout{
		grid:          grid,
		RowSizes:      make([]float64, rows),
		ColumnSizes:   make([]float64, columns),
		rowOffsets:    make([]float64, rows),
		columnOffsets: make([]float64, columns),
	}
	gridLayout.Recompute()
	return gridLayout
}

func (gl *GridLayout) Recompute() {
	for i := range gl.RowSizes {
		gl.RowSizes[i] = 1/float64(len(gl.grid)) + gl.rowOffsets[i]
	}
	for i := range gl.ColumnSizes {
		gl.ColumnSizes[i] = 1/float64(len(gl.grid[0])) + gl.columnOffsets[i]
	}
}

func (gl *GridLayout) Resize(rows, columns int) {
	// Resize the grid
	newGrid := make([][]gui.INode, rows)
	for i := range newGrid {
		newGrid[i] = make([]gui.INode, columns)
	}

	// Copy values from the old grid to the new grid
	for i := range gl.grid {
		copy(newGrid[i], gl.grid[i])
	}

	// Resize row sizes
	if rows >= len(gl.grid) {
		gl.RowSizes = append(gl.RowSizes, make([]float64, rows-len(gl.grid))...)
		gl.rowOffsets = append(gl.rowOffsets, make([]float64, rows-len(gl.grid))...)
	} else {
		gl.RowSizes = gl.RowSizes[:rows]
		gl.rowOffsets = gl.rowOffsets[:rows]
	}

	// Resize column sizes
	if columns >= len(gl.grid[0]) {
		gl.ColumnSizes = append(gl.ColumnSizes, make([]float64, columns-len(gl.grid[0]))...)
		gl.columnOffsets = append(gl.columnOffsets, make([]float64, columns-len(gl.grid[0]))...)
	} else {
		gl.ColumnSizes = gl.ColumnSizes[:columns]
		gl.columnOffsets = gl.columnOffsets[:columns]
	}

	gl.grid = newGrid

	// Recompute size percentages
	gl.Recompute()
}

func (gl *GridLayout) ResizeRow(row int, newSize float64) {
	if row > len(gl.grid)-2 {
		panic("Resized row does not exist or is not resizable (ie. last row)")
	}
	if newSize > 1 {
		newSize = 1
	}
	for i := range gl.RowSizes {
		gl.rowOffsets[i] -= newSize / float64(len(gl.grid)-1)
	}
	gl.rowOffsets[row] = newSize
	gl.Recompute()
}

func (gl *GridLayout) ResizeColumn(column int, newSize float64) {
	if column > len(gl.grid[0])-2 {
		panic("Resized column does not exist or is not resizable (ie. last column)")
	}
	if newSize > 1 {
		newSize = 1
	}
	for i := range gl.columnOffsets {
		gl.columnOffsets[i] -= newSize / float64(len(gl.grid)-1)
	}
	gl.columnOffsets[column] = newSize
	gl.Recompute()
}

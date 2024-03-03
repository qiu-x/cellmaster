package layout

import (
    "fmt"
)

type GridLayout struct {
    Rows         int
    Columns      int
    Grid         [][]int
    RowSizes     []float64
    ColumnSizes  []float64
}

func NewGridLayout(rows, columns int) *GridLayout {
    grid := make([][]int, rows)
    for i := range grid {
        grid[i] = make([]int, columns)
    }
    return &GridLayout{
        Rows:         rows,
        Columns:      columns,
        Grid:         grid,
        RowSizes:     make([]float64, rows),
        ColumnSizes:  make([]float64, columns),
    }
}

func (gl *GridLayout) SetCellValue(row, column, value int) {
    if row >= 0 && row < gl.Rows && column >= 0 && column < gl.Columns {
        gl.Grid[row][column] = value
    } else {
        fmt.Println("Error: Out of range")
    }
}

func (gl *GridLayout) GetCellValue(row, column int) int {
    if row >= 0 && row < gl.Rows && column >= 0 && column < gl.Columns {
        return gl.Grid[row][column]
    } else {
        fmt.Println("Error: Out of range")
        return -1 // or any default value indicating error
    }
}

func (gl *GridLayout) Resize(rows, columns int) {
    // Resize the grid
    newGrid := make([][]int, rows)
    for i := range newGrid {
        newGrid[i] = make([]int, columns)
    }

    // Copy values from the old grid to the new grid
    for i := 0; i < rows && i < gl.Rows; i++ {
        copy(newGrid[i], gl.Grid[i])
    }

    // Resize row sizes
    if rows >= gl.Rows {
        gl.RowSizes = append(gl.RowSizes, make([]float64, rows-gl.Rows)...)
    } else {
        gl.RowSizes = gl.RowSizes[:rows]
    }

    // Resize column sizes
    if columns >= gl.Columns {
        gl.ColumnSizes = append(gl.ColumnSizes, make([]float64, columns-gl.Columns)...)
    } else {
        gl.ColumnSizes = gl.ColumnSizes[:columns]
    }

    gl.Rows = rows
    gl.Columns = columns
    gl.Grid = newGrid

    // Recompute size percentages
    gl.Recompute()
}

func (gl *GridLayout) Recompute() {
    totalRowSizes := 0.0
    totalColumnSizes := 0.0

    for _, size := range gl.RowSizes {
        totalRowSizes += size
    }

    for _, size := range gl.ColumnSizes {
        totalColumnSizes += size
    }

    // Normalize row and column sizes to percentages
    for i := range gl.RowSizes {
        gl.RowSizes[i] = gl.RowSizes[i] / totalRowSizes * 100
    }

    for i := range gl.ColumnSizes {
        gl.ColumnSizes[i] = gl.ColumnSizes[i] / totalColumnSizes * 100
    }
}

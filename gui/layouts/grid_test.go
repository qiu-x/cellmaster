package layouts_test

import (
	"cellmaster/gui/layouts"
	"testing"
)

func TestNewGird(t *testing.T) {
	testGrid := layouts.NewGridLayout(4, 4)

	wantRows := []float64{0.25, 0.25, 0.25, 0.25}
	wantCols := []float64{0.25, 0.25, 0.25, 0.25}
	for i, got := range testGrid.RowSizes {
		if got != wantRows[i] {
			t.Errorf("got %f, wanted %f", got, wantRows[i])
		}
	}
	for i, got := range testGrid.ColumnSizes {
		if got != wantCols[i] {
			t.Errorf("got %f, wanted %f", got, wantCols[i])
		}
	}
}

func TestGirdResize(t *testing.T) {
	testGrid := layouts.NewGridLayout(3, 3)
	testGrid.Resize(4, 4)

	wantRows := []float64{0.25, 0.25, 0.25, 0.25}
	wantCols := []float64{0.25, 0.25, 0.25, 0.25}
	for i, got := range testGrid.RowSizes {
		if got != wantRows[i] {
			t.Errorf("got %f, wanted %f", got, wantRows[i])
		}
	}
	for i, got := range testGrid.ColumnSizes {
		if got != wantCols[i] {
			t.Errorf("got %f, wanted %f", got, wantCols[i])
		}
	}
}

func TestGirdRowResize(t *testing.T) {
	testGrid := layouts.NewGridLayout(3, 3)
	testGrid.Resize(4, 3)
	testGrid.ResizeRow(0, 0.15)

	wantRows := []float64{0.4, 0.2, 0.2, 0.2}
	for i, got := range testGrid.RowSizes {
		if got != wantRows[i] {
			t.Errorf("got %f, wanted %f", got, wantRows[i])
		}
	}
}

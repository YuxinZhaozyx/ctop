package process

import (
	ui "github.com/gizak/termui"
)

type ProcessGrid struct {
	ui.GridBufferer
	header *ProcessHeader
	cols   []ProcessCol
	Rows   []RowBufferer
	X, Y   int
	Width  int
	Height int
}

func NewProcessGrid() *ProcessGrid {
	pg := &ProcessGrid{
		header: NewProcessHeader(),
	}
	pg.rebuildHeader()
	return pg
}

func (pg *ProcessGrid) Align() {
	y := pg.Y

	colWidths := pg.calcWidths()
	for _, r := range pg.pageRows() {
		r.SetY(y)
		y += r.GetHeight()
		r.SetWidths(pg.Width, colWidths)
	}
}

func (pg *ProcessGrid) Clear() {
	pg.Rows = []RowBufferer{}
	pg.rebuildHeader()
}

func (pg *ProcessGrid) GetHeight() int {
	return len(pg.Rows) + pg.header.Height
}

func (pg *ProcessGrid) SetX(x int) {
	pg.X = x
}

func (pg *ProcessGrid) SetY(y int) {
	pg.Y = y
}

func (pg *ProcessGrid) SetWidth(w int) { pg.Width = w }

func (pg *ProcessGrid) MaxRows() int {
	return ui.TermHeight() - pg.header.Height - pg.Y
}

func (pg *ProcessGrid) calcWidths() []int {
	var autoCols int
	width := pg.Width
	colWidths := make([]int, len(pg.cols))

	for n, w := range pg.cols {
		colWidths[n] = w.FixedWidth()
		width -= w.FixedWidth()
		if w.FixedWidth() == 0 {
			autoCols++
		}
	}

	spacing := colSpacing * len(pg.cols)
	if autoCols > 0 {
		autoWidth := (width - spacing) / autoCols
		for n, val := range colWidths {
			if val == 0 {
				colWidths[n] = autoWidth
			}
		}
	}

	return colWidths
}

func (pg *ProcessGrid) pageRows() (rows []RowBufferer) {
	rows = append(rows, pg.header)
	rows = append(rows, pg.Rows...)
	return rows
}

func (pg *ProcessGrid) Buffer() ui.Buffer {
	buf := ui.NewBuffer()
	for _, r := range pg.pageRows() {
		buf.Merge(r.Buffer())
	}
	return buf
}

func (pg *ProcessGrid) AddRows(rows ...RowBufferer) {
	pg.Rows = append(pg.Rows, rows...)
}

func (pg *ProcessGrid) rebuildHeader() {
	pg.cols = newRowWidgets()
	pg.header.clearFieldPars()
	for _, col := range pg.cols {
		pg.header.addFieldPar(col.Header())
	}
}

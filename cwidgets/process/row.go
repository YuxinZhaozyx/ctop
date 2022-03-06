package process

import (
	ui "github.com/gizak/termui"
)

const rowPadding = 1


type RowBufferer interface {
	SetY(int)
	SetWidths(int, []int)
	GetHeight() int
	Buffer() ui.Buffer
}

type ProcessRow struct {
	Bg   *RowBg
	Cols []ProcessCol
	X, Y int
	Height int
	widths []int
}

func NewProcessRow() *ProcessRow {
	row := &ProcessRow{
		Bg:     NewRowBg(),
		Cols:   newRowWidgets(),
		X:      rowPadding,
		Height: 1,
	}
	return row
}

func (row *ProcessRow) SetMeta(m Meta) {
	for _, w := range row.Cols {
		w.SetMeta(m)
	}
}

func (row *ProcessRow) Reset() {
	for _, w := range row.Cols {
		w.Reset()
	}
}

func (row *ProcessRow) GetHeight() int {
	return row.Height
}

func (row *ProcessRow) SetY(y int) {
	if y == row.Y {
		return
	}

	row.Bg.Y = y
	for _, w := range row.Cols {
		w.SetY(y)
	}
	row.Y = y
}

func (row *ProcessRow) SetWidths(totalWidth int, widths []int) {
	x := row.X

	row.Bg.SetX(x)
	row.Bg.SetWidth(totalWidth)

	for n, w := range row.Cols {
		w.SetX(x)
		w.SetWidth(widths[n])
		x += widths[n] + colSpacing
	}
}

func (row *ProcessRow) Buffer() ui.Buffer {
	buf := ui.NewBuffer()
	buf.Merge(row.Bg.Buffer())
	for _, w := range row.Cols {
		buf.Merge(w.Buffer())
	}
	return buf
}

func (row *ProcessRow) Highlight() {
	for _, w := range row.Cols {
		w.Highlight()
	}
}

func (row *ProcessRow) UnHighlight() {
	for _, w := range row.Cols {
		w.UnHighlight()
	}
}

type RowBg struct {
	*ui.Par
}

func NewRowBg() *RowBg {
	bg := ui.NewPar("")
	bg.Height = 1
	bg.Border = false
	bg.Bg = ui.ThemeAttr("par.text.bg")
	return &RowBg{bg}
}

func (w *RowBg) Highlight()   { w.Bg = ui.ThemeAttr("par.text.fg") }
func (w *RowBg) UnHighlight() { w.Bg = ui.ThemeAttr("par.text.bg") }

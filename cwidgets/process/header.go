package process

import (
	ui "github.com/gizak/termui"
)

type ProcessHeader struct {
	X, Y   int
	Width  int
	Height int
	cols   []ProcessCol
	widths []int
	pars   []*ui.Par
}

func NewProcessHeader() *ProcessHeader {
	return &ProcessHeader{
		X:    rowPadding,
		Height: 2,
	}
}

func (row *ProcessHeader) GetHeight() int {
	return row.Height
}

func (row *ProcessHeader) SetWidths(totalWidth int, widths []int) {
	x := row.X

	for n, w := range row.pars {
		w.SetX(x)
		w.SetWidth(widths[n])
		x += widths[n] + colSpacing
	}
	row.Width = totalWidth
}

func (row *ProcessHeader) SetX(x int) {
	row.X = x
}

func (row *ProcessHeader) SetY(y int) {
	for _, p := range row.pars {
		p.SetY(y)
	}
	row.Y = y
}

func (row *ProcessHeader) Buffer() ui.Buffer {
	buf := ui.NewBuffer()
	for _, p := range row.pars {
		buf.Merge(p.Buffer())
	}
	return buf
}

func (row *ProcessHeader) clearFieldPars() {
	row.pars = []*ui.Par{}
}

func (row *ProcessHeader) addFieldPar(s string) {
	p := ui.NewPar(s)
	p.Height = row.Height
	p.Border = false
	row.pars = append(row.pars, p)
}

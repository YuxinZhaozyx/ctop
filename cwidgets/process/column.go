package process

import (
	ui "github.com/gizak/termui"
)

const colSpacing = 1

var (
	allCols = map[string]NewProcessColFn{
		"pid":   NewPIDCol,
		"user":  NewUserCol,
		"name":  NewNameCol,
		"cpu":   NewCPUCol,
		"mem":   NewMEMCol,
		"vsz":   NewVSZCol,
		"rss":   NewRSSCol,
		"start": NewStartCol,
		"command": NewCommandCol,
	}
)

type NewProcessColFn func() ProcessCol

func newRowWidgets() []ProcessCol {
	enabled := EnabledColumns()
	cols := make([]ProcessCol, len(enabled))

	for n, name := range enabled {
		wFn, ok := allCols[name]
		if !ok {
			panic("no such widget name: %s" + name)
		}
		cols[n] = wFn()
	}

	return cols
}

type ProcessCol interface {
	ui.GridBufferer
	Reset()
	Header() string
	FixedWidth() int
	Highlight()
	UnHighlight()
	SetMeta(Meta)
}

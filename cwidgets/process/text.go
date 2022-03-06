package process

import (
	"strings"

	ui "github.com/gizak/termui"
)

type MetaCol struct {
	*TextCol
	metaName string
}

func (w *MetaCol) SetMeta(m Meta) {
	w.setText(m.Get(w.metaName))
}

func NewPIDCol() ProcessCol {
	w := &MetaCol{NewTextCol("PID"), "pid"}
	w.fWidth = 4
	return w
}

func NewUserCol() ProcessCol {
	w := &MetaCol{NewTextCol("USER"), "user"}
	w.fWidth = 12
	return w
}

type NameCol struct {
	*TextCol
}

func NewNameCol() ProcessCol {
	w := &NameCol{NewTextCol("CONTAINER")}
	return w
}

func (w *NameCol) SetMeta(m Meta) {
	w.setText(strings.TrimPrefix(m["name"], "/"))
}

func NewVSZCol() ProcessCol {
	w := &MetaCol{NewTextCol("VSZ(KB)"), "vsz"}
	w.fWidth = 10
	return w

}

func NewRSSCol() ProcessCol {
	w := &MetaCol{NewTextCol("RSS(KB)"), "rss"}
	w.fWidth = 10
	return w
}

func NewCPUCol() ProcessCol {
	w := &MetaCol{NewTextCol("%CPU"), "cpu"}
	w.fWidth = 6
	return w
}

func NewMEMCol() ProcessCol {
	w := &MetaCol{NewTextCol("%MEM"), "mem"}
	w.fWidth = 6
	return w
}

func NewStartCol() ProcessCol {
	w := &MetaCol{NewTextCol("START"), "start"}
	w.fWidth = 8
	return w
}

type CommandCol struct {
	*TextCol
}

func NewCommandCol() ProcessCol {
	w := &CommandCol{NewTextCol("COMMAND")}
	return w
}

func (w *CommandCol) SetMeta(m Meta) {
	command_split := strings.Split(strings.Split(m["command"], " ")[0], "/")
	command_name := command_split[len(command_split) - 1]
	w.setText(command_name)
}


type TextCol struct {
	*ui.Par
	header string
	fWidth int
}

func NewTextCol(header string) *TextCol {
	p := ui.NewPar("-")
	p.Border = false
	p.Height = 1
	p.Width = 20

	return &TextCol{
		Par:    p,
		header: header,
		fWidth: 0,
	}
}

func (w *TextCol) Highlight() {
	w.Bg = ui.ThemeAttr("par.text.fg")
	w.TextFgColor = ui.ThemeAttr("par.text.hi")
	w.TextBgColor = ui.ThemeAttr("par.text.fg")
}

func (w *TextCol) UnHighlight() {
	w.Bg = ui.ThemeAttr("par.text.bg")
	w.TextFgColor = ui.ThemeAttr("par.text.fg")
	w.TextBgColor = ui.ThemeAttr("par.text.bg")
}

// TextCol implements CompactCol
func (w *TextCol) Reset()                    { w.setText("-") }
func (w *TextCol) SetMeta(Meta)       {}

func (w *TextCol) Header() string            { return w.header }
func (w *TextCol) FixedWidth() int           { return w.fWidth }

func (w *TextCol) setText(s string) {
	if w.fWidth > 0 && len(s) > w.fWidth {
		s = s[0:w.fWidth]
	}
	w.Text = s
}

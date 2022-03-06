package process

import (
	"strings"
	"fmt"
	"os"
	"sync"
)

var (
	GlobalColumns  []*Column
	lock           sync.RWMutex
)

var defaultColumns = []Column{
	{
		Name: "pid",
		Label: "PID",
		Enabled: true,
	},
	{
		Name: "user",
		Label: "User",
		Enabled: true,
	},
	{
		Name: "name",
		Label: "Container Name",
		Enabled: true,
	},
	{
		Name: "vsz",
		Label: "VSZ",
		Enabled: true,
	},
	{
		Name: "rss",
		Label: "RSS",
		Enabled: true,
	},
	{
		Name: "cpu",
		Label: "%CPU",
		Enabled: true,
	},
	{
		Name: "mem",
		Label: "%MEM",
		Enabled: true,
	},
	{
		Name: "start",
		Label: "START",
		Enabled: true,
	},
	{
		Name: "command",
		Label: "COMMAND",
		Enabled: true,
	},
}

func Init() {
	for _, c := range defaultColumns {
		x := c
		GlobalColumns = append(GlobalColumns, &x)
		log.Infof("loaded default widget config [%s]: %t", quote(x.Name), x.Enabled)
	}
}

func quote(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}

// Return env var value if set, else return defaultVal
func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return defaultVal
}


type Column struct {
	Name    string
	Label   string
	Enabled bool
}

func ColumnsString() string {
	return strings.Join(EnabledColumns(), ",")
}

func EnabledColumns() (a []string) {
	lock.RLock()
	defer lock.RUnlock()
	for _, col := range GlobalColumns {
		if col.Enabled {
			a = append(a, col.Name)
		}
	}
	return a
}


// ColumnToggle toggles the enabled status of a given column name
func ColumnToggle(name string) {
	col := GlobalColumns[colIndex(name)]
	col.Enabled = !col.Enabled
	log.Noticef("config change [column-%s]: %t -> %t", col.Name, !col.Enabled, col.Enabled)
}

// ColumnLeft moves the column with given name up one position, if possible
func ColumnLeft(name string) {
	idx := colIndex(name)
	if idx > 0 {
		swapCols(idx, idx-1)
	}
}

// ColumnRight moves the column with given name up one position, if possible
func ColumnRight(name string) {
	idx := colIndex(name)
	if idx < len(GlobalColumns)-1 {
		swapCols(idx, idx+1)
	}
}

// Set Column order and enabled status from one or more provided Column names
func SetColumns(names []string) {
	var (
		n          int
		curColStr  = ColumnsString()
		newColumns = make([]*Column, len(GlobalColumns))
	)

	lock.Lock()

	// add enabled columns by name
	for _, name := range names {
		newColumns[n] = popColumn(name)
		newColumns[n].Enabled = true
		n++
	}

	// extend with omitted columns as disabled
	for _, col := range GlobalColumns {
		newColumns[n] = col
		newColumns[n].Enabled = false
		n++
	}

	GlobalColumns = newColumns
	lock.Unlock()

	log.Noticef("config change [columns]: %s -> %s", curColStr, ColumnsString())
}

func swapCols(i, j int) { GlobalColumns[i], GlobalColumns[j] = GlobalColumns[j], GlobalColumns[i] }

func popColumn(name string) *Column {
	idx := colIndex(name)
	if idx < 0 {
		panic("no such column name: " + name)
	}
	col := GlobalColumns[idx]
	GlobalColumns = append(GlobalColumns[:idx], GlobalColumns[idx+1:]...)
	return col
}

// return index of column with given name, if any
func colIndex(name string) int {
	for n, c := range GlobalColumns {
		if c.Name == name {
			return n
		}
	}
	return -1
}

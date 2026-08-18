package modules

//go:generate stringer -type=Module
type Module int

const (
	Amazon Module = iota
)

var LENGTH Module = Module(len(_Module_index) - 1)

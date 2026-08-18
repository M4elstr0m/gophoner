package modules

import "github.com/charmbracelet/log"

const defaultModuleColor = "#000000"

var colorRegistry = map[Module]string{}

func RegisterColor(m Module, color string) {
	_, exists := colorRegistry[m]
	if exists {
		log.Fatal("Attempt to register module color multiple times",
			"module", m,
		)
	}

	colorRegistry[m] = color
}

func (m Module) Color() string {
	str, ok := colorRegistry[m]
	if !ok {
		str = defaultModuleColor
	}

	return str
}

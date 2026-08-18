package modules

import (
	"fmt"
	"strings"
)

var byName = func() map[string]Module {
	m := make(map[string]Module, LENGTH)
	for i := Module(0); i < LENGTH; i++ {
		m[strings.ToLower(i.String())] = i
	}
	return m
}()

func Parse(name string) (Module, error) {
	m, ok := byName[strings.ToLower(strings.TrimSpace(name))]
	if !ok {
		return 0, fmt.Errorf("unknown module %q", name)
	}
	return m, nil
}

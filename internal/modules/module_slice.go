package modules

import "strings"

type ModuleSlice []Module

func (l *ModuleSlice) String() string {
	names := make([]string, len(*l))
	for i, m := range *l {
		names[i] = m.String()
	}
	return strings.Join(names, ",")
}

func (l *ModuleSlice) Set(value string) error {
	for _, part := range strings.Split(value, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		m, err := Parse(part)
		if err != nil {
			return err
		}
		*l = append(*l, m)
	}
	return nil
}

func (l *ModuleSlice) Type() string {
	return "modules"
}

func All() ModuleSlice {
	all := make(ModuleSlice, 0, LENGTH)
	for m := Module(0); m < LENGTH; m++ {
		if _, ok := Get(m); ok {
			all = append(all, m)
		}
	}
	return all
}

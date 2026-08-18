package tui

import (
	"github.com/M4elstr0m/gophoner/internal/modules"
	"github.com/M4elstr0m/gophoner/internal/report"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type modulesSubmittedMsg struct {
	Modules modules.ModuleSlice
}

type backMsg struct{}

type moduleItem struct {
	module  modules.Module
	checked bool
}

type moduleView struct {
	width, height int

	items   []moduleItem
	focused int

	err string
}

var (
	moduleRowStyle = lipgloss.NewStyle().Padding(0, 1)

	moduleRowFocusedStyle = moduleRowStyle.
				Background(lipgloss.Color("#7D56F4")).
				Foreground(lipgloss.Color("#FFFFFF"))
)

const (
	checkedGlyph   = "[x]"
	uncheckedGlyph = "[ ]"
)

func newModuleView() *moduleView {
	all := modules.All()
	items := make([]moduleItem, len(all))
	for i, m := range all {
		items[i] = moduleItem{module: m}
	}

	return &moduleView{items: items}
}

func (v *moduleView) proceedIndex() int {
	return len(v.items) + 1
}

func (v *moduleView) cancelIndex() int {
	return len(v.items) + 2
}

func (v *moduleView) totalStops() int {
	return len(v.items) + 3
}

func (v *moduleView) Init() tea.Cmd {
	return nil
}

func (v *moduleView) Update(msg tea.Msg) (View, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		v.width, v.height = msg.Width, msg.Height
		return v, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return v, func() tea.Msg { return backMsg{} }

		case "up", "k":
			v.focused = (v.focused - 1 + v.totalStops()) % v.totalStops()
			return v, nil

		case "down", "j", "tab":
			v.focused = (v.focused + 1) % v.totalStops()
			return v, nil

		case "left":
			if v.focused == v.cancelIndex() {
				v.focused = v.proceedIndex()
			}
			return v, nil

		case "right":
			if v.focused == v.proceedIndex() {
				v.focused = v.cancelIndex()
			}
			return v, nil

		case " ":
			v.toggle()
			return v, nil

		case "enter":
			switch v.focused {
			case v.proceedIndex():
				return v.submit()
			case v.cancelIndex():
				return v, func() tea.Msg { return backMsg{} }
			default:
				v.toggle()
			}
			return v, nil
		}
	}

	return v, nil
}

func (v *moduleView) toggle() {
	switch {
	case v.focused == 0:
		v.setAll(!v.allChecked())
	case v.focused < v.proceedIndex():
		v.items[v.focused-1].checked = !v.items[v.focused-1].checked
	}
}

func (v *moduleView) setAll(state bool) {
	for i := range v.items {
		v.items[i].checked = state
	}
}

func (v *moduleView) allChecked() bool {
	if len(v.items) == 0 {
		return false
	}
	for _, item := range v.items {
		if !item.checked {
			return false
		}
	}
	return true
}

func (v *moduleView) submit() (View, tea.Cmd) {
	selected := make(modules.ModuleSlice, 0, len(v.items))
	for _, item := range v.items {
		if item.checked {
			selected = append(selected, item.module)
		}
	}

	if len(selected) == 0 {
		v.err = "Select at least one module"
		return v, nil
	}

	v.err = ""
	return v, func() tea.Msg {
		return modulesSubmittedMsg{Modules: selected}
	}
}

func (v *moduleView) View() string {
	rows := make([]string, 0, len(v.items)+3)

	rows = append(rows, labelStyle.Render("Select modules"))
	rows = append(rows, v.row(0, checkboxGlyph(v.allChecked())+" Select All"))
	rows = append(rows, "")

	for i, item := range v.items {
		label := lipgloss.JoinHorizontal(lipgloss.Top,
			checkboxGlyph(item.checked), " ", report.ModuleBadge(item.module))
		rows = append(rows, v.row(i+1, label))
	}

	proceed := buttonStyle
	cancel := buttonStyle
	if v.focused == v.proceedIndex() {
		proceed = buttonFocusedStyle
	}
	if v.focused == v.cancelIndex() {
		cancel = buttonFocusedStyle
	}
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, proceed.Render("PROCEED"), "  ", cancel.Render("CANCEL"))
	rows = append(rows, "", buttons)

	if v.err != "" {
		rows = append(rows, errorStyle.Render(v.err))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, rows...)
	content = padForEvenCentering(content, v.width)

	return lipgloss.Place(v.width, v.height, lipgloss.Center, lipgloss.Center, content)
}

func (v *moduleView) row(index int, label string) string {
	style := moduleRowStyle
	if v.focused == index {
		style = moduleRowFocusedStyle
	}
	return style.Render(label)
}

func checkboxGlyph(checked bool) string {
	if checked {
		return checkedGlyph
	}
	return uncheckedGlyph
}

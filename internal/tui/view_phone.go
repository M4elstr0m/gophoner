package tui

import (
	"strings"

	_countries "github.com/M4elstr0m/gophoner/internal/utils/countries"
	_strings "github.com/M4elstr0m/gophoner/internal/utils/strings"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type phoneSubmittedMsg struct {
	CountryCode string
	PhoneNumber string
}

type cancelledMsg struct{}

type phoneFocus int

const (
	focusCountryCode phoneFocus = iota
	focusPhoneNumber
	focusNext
	focusCancel
)

const (
	countryFieldWidth = 8
	phoneFieldWidth   = 22
	countryBadgeGap   = "  "
)

type phoneView struct {
	countryCode textinput.Model
	phoneNumber textinput.Model
	focused     phoneFocus

	err string
}

var (
	countryFieldBoxStyle = lipgloss.NewStyle().Width(countryFieldWidth)
	phoneFieldBoxStyle   = lipgloss.NewStyle().Width(phoneFieldWidth)

	countryBadgeStyle = lipgloss.NewStyle().
				Bold(true).
				Padding(0, 1).
				Foreground(lipgloss.Color("#FFFFFF"))

	countryBadgeSlotWidth = lipgloss.Width(countryBadgeGap + countryBadgeStyle.Render("XX"))

	countryBadgeDefaultColor = "#444444"
)

func newPhoneView() *phoneView {
	countryCode := textinput.New()
	countryCode.Placeholder = "+1"
	countryCode.CharLimit = 4
	countryCode.Width = 6
	countryCode.Validate = _strings.DigitsWithOptionalPlus
	countryCode.Focus()

	phoneNumber := textinput.New()
	phoneNumber.Placeholder = "5551234567"
	phoneNumber.CharLimit = 15
	phoneNumber.Width = 20
	phoneNumber.Validate = _strings.DigitsOnly

	return &phoneView{
		countryCode: countryCode,
		phoneNumber: phoneNumber,
		focused:     focusCountryCode,
	}
}

func newPhoneViewWithValues(countryCode, phoneNumber string) *phoneView {
	v := newPhoneView()

	v.countryCode.SetValue(countryCode)
	v.countryCode.CursorEnd()

	v.phoneNumber.SetValue(phoneNumber)
	v.phoneNumber.CursorEnd()

	return v
}

func (v *phoneView) Init() tea.Cmd {
	return textinput.Blink
}

func (v *phoneView) Update(msg tea.Msg) (View, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return v, func() tea.Msg { return cancelledMsg{} }

		case "tab", "down":
			v.setFocus((v.focused + 1) % 4)
			return v, nil

		case "shift+tab", "up":
			v.setFocus((v.focused + 3) % 4)
			return v, nil

		case "left":
			if v.focused == focusCancel {
				v.setFocus(focusNext)
			}
			return v, nil

		case "right":
			if v.focused == focusNext {
				v.setFocus(focusCancel)
			}
			return v, nil

		case "enter":
			switch v.focused {
			case focusCountryCode, focusPhoneNumber:
				v.setFocus(v.focused + 1)
				return v, nil
			case focusNext:
				return v.submit()
			case focusCancel:
				return v, func() tea.Msg { return cancelledMsg{} }
			}
		}
	}

	var cmd tea.Cmd
	switch v.focused {
	case focusCountryCode:
		v.countryCode, cmd = v.countryCode.Update(msg)
	case focusPhoneNumber:
		v.phoneNumber, cmd = v.phoneNumber.Update(msg)
	}

	return v, cmd
}

func (v *phoneView) setFocus(f phoneFocus) {
	v.focused = f

	if f == focusCountryCode {
		v.countryCode.Focus()
	} else {
		v.countryCode.Blur()
	}

	if f == focusPhoneNumber {
		v.phoneNumber.Focus()
	} else {
		v.phoneNumber.Blur()
	}
}

func (v *phoneView) submit() (View, tea.Cmd) {
	code := strings.TrimPrefix(strings.TrimSpace(v.countryCode.Value()), "+")
	number := strings.TrimSpace(v.phoneNumber.Value())

	if code == "" {
		v.err = "Enter a valid country code"
		v.setFocus(focusCountryCode)
		return v, nil
	}

	if number == "" {
		v.err = "Enter a phone number"
		v.setFocus(focusPhoneNumber)
		return v, nil
	}

	v.err = ""
	return v, func() tea.Msg {
		return phoneSubmittedMsg{CountryCode: "+" + code, PhoneNumber: number}
	}
}

func (v *phoneView) View() string {
	badge := ""
	if iso, ok := _countries.CountryCodeToISO(v.countryCode.Value()); ok {
		color, ok := _countries.ColorForISO(iso)
		if !ok {
			color = countryBadgeDefaultColor
		}
		badgeStyle := countryBadgeStyle.Background(lipgloss.Color(color))
		badge = countryBadgeGap + badgeStyle.Render(iso)
	}
	badgeSlot := lipgloss.NewStyle().Width(countryBadgeSlotWidth).Render(badge)

	countryLine := lipgloss.JoinHorizontal(lipgloss.Top,
		countryFieldBoxStyle.Render(v.countryCode.View()),
		badgeSlot,
	)

	phoneLine := phoneFieldBoxStyle.Render(v.phoneNumber.View())

	next := buttonStyle
	cancel := buttonStyle
	if v.focused == focusNext {
		next = buttonFocusedStyle
	}
	if v.focused == focusCancel {
		cancel = buttonFocusedStyle
	}
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, next.Render("NEXT"), "  ", cancel.Render("QUIT"))

	sections := []string{labelStyle.Render("Enter a phone number"), countryLine, phoneLine}
	if v.err != "" {
		sections = append(sections, errorStyle.Render(v.err))
	}
	sections = append(sections, "", buttons)

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

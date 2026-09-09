package report

import (
	"fmt"

	"github.com/M4elstr0m/gophoner/internal/recon"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const progressBarWidth = 40

const (
	moduleStartShare  = 0.1
	moduleFinishShare = 0.9
)

type moduleEventMsg struct {
	event recon.Event
	ok    bool
}

type progressModel struct {
	bar     progress.Model
	channel <-chan recon.Event

	total          int
	perModuleShare float64
	percent        float64

	results []recon.Output
	done    bool
}

func newProgressModel(input *recon.Input) *progressModel {
	total := len(input.ModuleSlice)

	return &progressModel{
		bar:            progress.New(progress.WithDefaultGradient(), progress.WithWidth(progressBarWidth)),
		channel:        recon.Run(input),
		total:          total,
		perModuleShare: 1 / float64(total),
		results:        make([]recon.Output, 0, total),
	}
}

func waitForModuleEvent(ch <-chan recon.Event) tea.Cmd {
	return func() tea.Msg {
		event, ok := <-ch
		return moduleEventMsg{event: event, ok: ok}
	}
}

func (m *progressModel) Init() tea.Cmd {
	return waitForModuleEvent(m.channel)
}

func (m *progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case moduleEventMsg:
		if !msg.ok {
			m.done = true
			if !m.bar.IsAnimating() {
				return m, tea.Quit
			}
			return m, nil
		}

		switch msg.event.Kind {
		case recon.ModuleStarted:
			m.percent += m.perModuleShare * moduleStartShare
		case recon.ModuleFinished:
			m.percent += m.perModuleShare * moduleFinishShare
			m.results = append(m.results, msg.event.Output)
		}
		if m.percent > 1 {
			m.percent = 1
		}

		percentCmd := m.bar.SetPercent(m.percent)
		return m, tea.Batch(percentCmd, waitForModuleEvent(m.channel))

	case progress.FrameMsg:
		newBar, cmd := m.bar.Update(msg)
		m.bar = newBar.(progress.Model)
		if m.done && !m.bar.IsAnimating() {
			return m, tea.Quit
		}
		return m, cmd
	}

	return m, nil
}

func (m *progressModel) View() string {
	return fmt.Sprintf("Checking modules %s %d/%d\n\n", m.bar.View(), len(m.results), m.total)
}

func Run(input *recon.Input) ([]recon.Output, error) {
	finalModel, err := tea.NewProgram(newProgressModel(input)).Run()
	if err != nil {
		return nil, err
	}

	return finalModel.(*progressModel).results, nil
}

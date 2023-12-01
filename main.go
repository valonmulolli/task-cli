package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const divisor = 4

const (
	todo status = iota
	inProgress
	done
)

type Task struct {
	status      status
	title       string
	description string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) FilterField() string {
	return "title"
}

func (t Task) Description() string {
	return t.description
}

type Model struct {
	loaded  bool
	focused status
	lists   []list.Model
	err     error
}

func New() *Model {
	return &Model{}
}

func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}
	// init Tasks
	m.lists[todo].Title = "Tasks"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "learn Go", description: "Go Programming"},
		Task{status: todo, title: "learn Rust", description: "Rust Programming"},
		Task{status: todo, title: "Learn Zig", description: "Zig Programming"},
	})
	// init in progress
	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{status: todo, title: "expoloring Go", description: "Go Programming"},
	})
	// init done
	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{status: todo, title: "Finally i learned Go", description: "Go Programming"},
	})

}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.loaded {
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.lists[todo].View(),
			m.lists[inProgress].View(),
			m.lists[done].View(),
		)
	} else {
		return "Loading..."
	}
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

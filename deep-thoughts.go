package main

import (
	"deep-thoughts/thoughts"
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	log.SetFlags(0)

	p := tea.NewProgram(model{})

	m, err := p.Run()
	if err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}

	if m, ok := m.(model); ok && m.choice != "" {

		switch m.choice {
		case choices[0]:
			thoughts.Vote(m.deepThought.Id, true)
			log.Println("\n---\n👍 Upvoted.")
		case choices[1]:
			thoughts.Vote(m.deepThought.Id, false)
			log.Println("\n---\n👎 Downvoted.")
		case choices[2]:
			log.Println("\n---\nAdios!")
		}
	}
}

var choices = []string{"👍 Upvote", "👎 Downvote", "❌ Cancel"}

func (m model) Init() tea.Cmd {
	return getRandomThought
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}

	case thoughts.DeepThought:
		m.deepThought = msg
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}

	s.WriteString(fmt.Sprintf(`%s`, m.deepThought.Text))
	s.WriteString("\n\nWhat do you think?\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

type model struct {
	cursor      int
	choice      string
	deepThought thoughts.DeepThought
}

func getRandomThought() tea.Msg {
	thought, err := thoughts.Random()

	if err != nil {
		log.Fatalf(err.Error())
	}

	return thought
}

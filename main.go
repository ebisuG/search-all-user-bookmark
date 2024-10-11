package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	// choices  []string         // items on the to-do list
	// cursor   int              // which to-do list item our cursor is pointing at
	// selected map[int]struct{} // which to-do items are selected
	searchPath   []string
	searchString textinput.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Search keyword"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	username := os.Getenv("WindowsUserName")
	var searchPathString strings.Builder
	searchPathString.WriteString("C:\\Users\\")
	searchPathString.WriteString(username)
	searchPathString.WriteString("\\AppData\\Local\\Google\\Chrome\\User Data")

	return model{
		// Our to-do list is a grocery list
		// choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		// selected: make(map[int]struct{}),
		searchPath:   []string{searchPathString.String()},
		searchString: ti,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// fmt.Println("msg : ", msg)
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c":
			fmt.Println("Bye bye!")
			return m, tea.Quit

			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			fmt.Println("at the first part of enter")
			// reader := bufio.NewReader(os.Stdin)
			// text, _ := reader.ReadString('\n')
			// m.searchString = text
			fmt.Println("at the last part of enter")
			// _, ok := m.selected[m.cursor]
			// if ok {
			// 	delete(m.selected, m.cursor)
			// } else {
			// 	m.selected[m.cursor] = struct{}{}
			// }
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	m.searchString, cmd = m.searchString.Update(msg)
	return m, cmd
}

func (m model) View() string {
	// The header
	s := "Below string is being searched... \n\n"
	// s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	// for i, choice := range m.choices {

	// 	// Is the cursor pointing at this choice?
	// 	cursor := " " // no cursor
	// 	if m.cursor == i {
	// 		cursor = ">" // cursor!
	// 	}

	// 	// Is this choice selected?
	// 	checked := " " // not selected
	// 	if _, ok := m.selected[i]; ok {
	// 		checked = "x" // selected!
	// 	}

	// 	// Render the row
	// 	s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	// }
	s += fmt.Sprintf("%s\n", m.searchString.View())

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return fmt.Sprintf("%s\n", m.searchString.View())
	// return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
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
	fmt.Println("Press Ctrl + c when you quit")
	fmt.Println("searchPathString : ", searchPathString.String())

	return model{
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
			// fmt.Println("enter is pushed")
			// fmt.Println("searchString : ", m.searchString.View()[2:], "\nnewline")
			data, err := os.ReadFile(m.searchPath[0] + "\\Default\\Bookmarks")
			checkError(err)
			fmt.Println(string(data))
			fmt.Println("")
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
	s += fmt.Sprintf("%s\n", m.searchString.View())

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return fmt.Sprintf("%s\n", m.searchString.View())
	// return s
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/ebisuG/search-all-user-bookmark/internal/bubbleTea"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(bubbleTea.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

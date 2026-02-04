package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type InfoDisplayed struct {
	Name          string
	Url           string
	BookmarkTitle BookmarkTitle
	BookmarkUrl   BookmarkUrl
}

func (m model) View() string {
	s := "Below string is being searched... \n\n"
	s += fmt.Sprintf("%s\n", m.searchString.View())

	s += "\nPress q to quit.\n"

	return fmt.Sprintf("%s\n", m.searchString.View())
}

func FormatDisplay(info []InfoDisplayed) {
	nameColor := lipgloss.Color("#F77F0F")

	for _, v := range info {
		hypelink := termenv.Hyperlink(v.BookmarkUrl.Record.Raw, v.BookmarkTitle.Record.Raw)
		fmt.Println(" ", lipgloss.NewStyle().Foreground(nameColor).Bold(false).Render(hypelink))
	}
}

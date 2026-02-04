package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	searchPath   []string
	searchString textinput.Model
	allUrl       []InfoDisplayed
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Search keyword"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#a871f0")).Bold(true).SetString("Press Ctrl + c when you quit"))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#a871f0")).Bold(true).SetString("Press Ctrl to jummp to bookmark"))
	fmt.Println("Reading all bookmark files...")
	bookmarkFilesPath := GetAllBookmarkFilePath()
	var allData []InfoDisplayed
	defaultData, err := ReadBookmarkFile(GetPathName() + "\\Default\\Bookmarks")
	if err != nil {
		log.Fatal(err)
	}
	allData = append(allData, defaultData...)
	for _, v := range bookmarkFilesPath {
		profileData, err := ReadBookmarkFile(v)
		if err != nil {
			fmt.Println("No file : ", v)
			continue
		}
		allData = append(allData, profileData...)
	}
	fmt.Println("Finish reading bookmark files...")
	return model{
		searchPath:   []string{GetPathName()},
		searchString: ti,
		allUrl:       allData,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

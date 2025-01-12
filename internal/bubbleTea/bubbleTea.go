package bubbleTea

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ebisuG/search-all-user-bookmark/internal/util"
)

type model struct {
	searchPath   []string
	searchString textinput.Model
	allUrl       []util.InfoDisplayed
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Search keyword"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#a871f0")).Bold(true).SetString("Press Ctrl + c when you quit"))
	fmt.Println("Reading all bookmark files...")
	bookmarkFilesPath := util.GetAllBookmarkFilePath()
	var allData []util.InfoDisplayed
	defaultData, err := util.ReadBookmarkFile(util.GetPathName() + "\\Default\\Bookmarks")
	if err != nil {
		log.Fatal(err)
	}
	allData = append(allData, defaultData...)
	for _, v := range bookmarkFilesPath {
		profileData, err := util.ReadBookmarkFile(v)
		if err != nil {
			fmt.Println("No file : ", v)
			continue
		}
		allData = append(allData, profileData...)
	}
	fmt.Println("Finish reading bookmark files...")
	return model{
		searchPath:   []string{util.GetPathName()},
		searchString: ti,
		allUrl:       allData,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			fmt.Println("Bye bye!")
			return m, tea.Quit

		case "enter", " ":
			var searchWord []string
			for _, v := range m.searchString.Value() {
				searchWord = append(searchWord, string(v))
			}
			var display []util.InfoDisplayed
			display = util.FilterByString(m.allUrl, strings.Join(searchWord, ""))
			FormatDisplay(display)
			fmt.Println("")
		}
	}

	m.searchString, cmd = m.searchString.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := "Below string is being searched... \n\n"
	s += fmt.Sprintf("%s\n", m.searchString.View())

	s += "\nPress q to quit.\n"

	return fmt.Sprintf("%s\n", m.searchString.View())
}

func FormatDisplay(info []util.InfoDisplayed) {
	nameColor := lipgloss.Color("#F77F0F")
	urlColor := lipgloss.Color("#FAEECA")
	for _, v := range info {
		fmt.Println(lipgloss.NewStyle().Foreground(nameColor).Bold(true).SetString(v.Name),
			lipgloss.NewStyle().Foreground(urlColor).Bold(true).MarginBottom(1).SetString(v.Url),
		)
	}
}

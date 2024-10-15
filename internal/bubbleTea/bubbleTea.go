package bubbleTea

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
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

	fmt.Println("Press Ctrl + c when you quit")
	fmt.Println("Reading all bookmark files...")
	bookmarkFilesPath := util.GetAllBookmarkFilePath()
	var allData []util.InfoDisplayed
	allData = append(allData, util.ReadBookmarkFile(util.GetPathName()+"\\Default\\Bookmarks")...)
	for _, v := range bookmarkFilesPath {
		allData = append(allData, util.ReadBookmarkFile(v)...)
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

			for _, v := range display {
				fmt.Println(v.Name, " : ", v.Url, "")
			}
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

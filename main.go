package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ebisuG/search-all-user-bookmark/internal/config"
	"github.com/ebisuG/search-all-user-bookmark/internal/infra"
	"github.com/ebisuG/search-all-user-bookmark/internal/search"
	"github.com/muesli/termenv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	searchString textinput.Model
	config       config.Config
}

type hit struct {
	bookmarkTitle bookmarkTitle
	bookmarkUrl   bookmarkUrl
}

type result []hit

type bookmarkTitle struct {
	record record
}

type bookmarkUrl struct {
	record record
}

type record struct {
	raw  string
	norm string
}

func NewChromeLoader() infra.ChromeLoader {
	return infra.ChromeLoader{}
}
func NewChromeFinder() infra.ChromeFinder {
	return infra.ChromeFinder{}
}
func NewChromeParser() infra.ChromeParser {
	return infra.ChromeParser{}
}
func NewSearcher() search.Searcher {
	return infra.CoreSearcher{}
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

	var config config.Config
	chromeLoader := NewChromeLoader()
	chromeFinder := NewChromeFinder()
	clisetting, err := chromeLoader.Load("./settings.json")
	if err != nil {
		fmt.Println(err)
		return model{}
	}
	config.CliSetting = clisetting
	fmt.Println("Finish loading settings.json")
	config.SearchPath, err = chromeFinder.Find(config.CliSetting)

	return model{searchString: ti, config: config}
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
			var bookmarks search.Bookmarks
			chromeParser := NewChromeParser()
			for _, v := range m.config.SearchPath {
				bookmark, err := chromeParser.Parse(v)
				if err != nil {
					fmt.Println(v)
					fmt.Println("failed to parse chrome bookmark file")
				}
				bookmarks = append(bookmarks, bookmark...)
			}

			searcher := NewSearcher()
			display, err := searcher.Search(bookmarks, strings.Join(searchWord, ""))
			if err != nil {
				fmt.Println("failed to search bookmarks")
			}
			var result result
			for _, v := range display {
				titleRecord := record{norm: v.BookmarkTitle.Record.Norm, raw: v.BookmarkTitle.Record.Raw}
				urlRecord := record{norm: v.BookmarkUrl.Record.Norm, raw: v.BookmarkUrl.Record.Raw}
				result = append(result, hit{bookmarkTitle: bookmarkTitle{titleRecord}, bookmarkUrl: bookmarkUrl{urlRecord}})
			}
			result.FormatDisplay()
			m.searchString.Reset()
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

func (r result) FormatDisplay() {
	nameColor := lipgloss.Color("#F77F0F")

	for _, v := range r {
		hypelink := termenv.Hyperlink(v.bookmarkUrl.record.raw, v.bookmarkTitle.record.raw)
		fmt.Println(" ", lipgloss.NewStyle().Foreground(nameColor).Bold(false).Render(hypelink))
	}
}

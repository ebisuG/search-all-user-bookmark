package main

import (
	"encoding/json"
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

type ParentJson struct {
	Checksum     string `json:"checksum"`
	Roots        Roots  `json:"roots"`
	SyncMetadata string `json:"sync_metadata"`
	Version      int    `json:"version"`
}

type Roots struct {
	BookmarkBar BookmarkBar `json:"bookmark_bar"`
}

type BookmarkBar struct {
	Children []Child `json:"children"`
	Other    Child   `json:"other"`
	Synced   Child   `json:"synced"`
}

type Child struct {
	Children     []Child   `json:"children"`
	DateAdded    string    `json:"date_added"`
	DateLastUsed string    `json:"date_last_used"`
	Guid         string    `json:"guid"`
	Id           string    `json:"id"`
	MetaInfo     Meta_info `json:"meta_info"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Url          string    `json:"url"`
}

type Meta_info struct {
	PowerBookmarkMeta string `json:"power_bookmark_meta"`
}

type InfoDisplayed struct {
	name string
	url  string
}

type SettingFile struct {
	UserName string `json:"username"`
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Search keyword"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	fmt.Println("Press Ctrl + c when you quit")
	return model{
		searchPath:   []string{getPathName()},
		searchString: ti,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

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
			data, err := os.ReadFile(m.searchPath[0] + "\\Default\\Bookmarks")
			checkError(err)
			var bookmarks ParentJson
			var display []InfoDisplayed
			var searchWord []string
			for _, v := range m.searchString.Value() {
				searchWord = append(searchWord, string(v))
			}
			json.Unmarshal(data, &bookmarks)
			for i := 0; i < len(bookmarks.Roots.BookmarkBar.Children); i++ {
				bookmark := bookmarks.Roots.BookmarkBar.Children[i]
				display = append(display, getChildren(bookmark)...)
			}
			display = filterByString(display, strings.Join(searchWord, ""))

			for _, v := range display {
				fmt.Println(v.name, " : ", v.url, "")
			}
			fmt.Println("")
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	m.searchString, cmd = m.searchString.Update(msg)
	return m, cmd
}

func getChildren(c Child) []InfoDisplayed {
	var result []InfoDisplayed
	if c.Type == "folder" {
		for _, v := range c.Children {
			result = append(result, getChildren(v)...)
		}
	} else {
		var pair InfoDisplayed
		pair.name, pair.url = c.Name, c.Url
		result = append(result, pair)
	}
	return result
}

func filterByString(pairs []InfoDisplayed, search string) []InfoDisplayed {
	var result []InfoDisplayed
	for _, v := range pairs {
		isInName := strings.Contains(v.name, search)
		isInUrl := strings.Contains(v.url, search)
		if isInName || isInUrl {
			result = append(result, v)
		}
	}
	return result
}

func getPathName() string {
	data, err := os.ReadFile("./settings.json")
	if err != nil {
		panic(err)
	}
	var settings SettingFile
	json.Unmarshal(data, &settings)
	var searchPathString strings.Builder
	searchPathString.WriteString("C:\\Users\\")
	searchPathString.WriteString(settings.UserName)
	searchPathString.WriteString("\\AppData\\Local\\Google\\Chrome\\User Data")
	return searchPathString.String()
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

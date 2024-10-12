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

//Handle file data as json
//get data
//search by key-map

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
			var bookmarks ParentJson
			json.Unmarshal(data, &bookmarks)
			// fmt.Println(bookmarks.Roots.BookmarkBar)
			// fmt.Println(bookmarks.Roots.BookmarkBar.Children)
			for i := 0; i < len(bookmarks.Roots.BookmarkBar.Children); i++ {
				bookmark := bookmarks.Roots.BookmarkBar.Children[i]
				fmt.Println(getChildren(bookmark))
				// fmt.Println(bookmark.Children)
				fmt.Println("----------")
			}
			// text, err := json.Marshal(data)
			// fmt.Println("text : ", text)

			// fmt.Println(string(data))
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

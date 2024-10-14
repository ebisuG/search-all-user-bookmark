package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	searchPath   []string
	searchString textinput.Model
	allUrl       []InfoDisplayed
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
	fmt.Println("Reading all bookmark files...")
	bookmarkFilesPath := getAllBookmarkFilePath()
	var allData []InfoDisplayed
	allData = append(allData, readBookmarkFile(getPathName()+"\\Default\\Bookmarks")...)
	for _, v := range bookmarkFilesPath {
		allData = append(allData, readBookmarkFile(v)...)
	}
	fmt.Println("Finish reading bookmark files...")
	return model{
		searchPath:   []string{getPathName()},
		searchString: ti,
		allUrl:       allData,
	}
}

func readBookmarkFile(path string) []InfoDisplayed {
	data, err := os.ReadFile(path)
	checkError(err)
	var bookmarks ParentJson
	json.Unmarshal(data, &bookmarks)
	var display []InfoDisplayed
	for i := 0; i < len(bookmarks.Roots.BookmarkBar.Children); i++ {
		bookmark := bookmarks.Roots.BookmarkBar.Children[i]
		display = append(display, getChildren(bookmark)...)
	}
	return display
}

func getAllBookmarkFilePath() []string {
	var bookmarksFilePath []string
	pathName := getPathName()
	files, err := os.ReadDir(pathName)
	if err != nil {
		panic(err)
	}
	r, _ := regexp.Compile("^Profile [0-9]*")

	for _, v := range files {
		match := r.MatchString(v.Name())
		if v.IsDir() && match {
			bookmarksFilePath = append(bookmarksFilePath, pathName+"\\"+v.Name()+"\\Bookmarks")
		}
	}
	return bookmarksFilePath
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
			var display []InfoDisplayed
			display = filterByString(m.allUrl, strings.Join(searchWord, ""))

			for _, v := range display {
				fmt.Println(v.name, " : ", v.url, "")
			}
			fmt.Println("")
		}
	}

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
	s := "Below string is being searched... \n\n"
	s += fmt.Sprintf("%s\n", m.searchString.View())

	s += "\nPress q to quit.\n"

	return fmt.Sprintf("%s\n", m.searchString.View())
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

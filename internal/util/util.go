package util

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type InfoDisplayed struct {
	Name string
	Url  string
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

type SettingFile struct {
	UserName string `json:"username"`
}

func ReadBookmarkFile(path string) ([]InfoDisplayed, error) {
	var display []InfoDisplayed
	data, err := os.ReadFile(path)
	if err != nil {
		return display, errors.New("no file")
	}
	var bookmarks ParentJson
	json.Unmarshal(data, &bookmarks)
	for i := 0; i < len(bookmarks.Roots.BookmarkBar.Children); i++ {
		bookmark := bookmarks.Roots.BookmarkBar.Children[i]
		display = append(display, GetChildren(bookmark)...)
	}
	return display, nil
}

func GetAllBookmarkFilePath() []string {
	var bookmarksFilePath []string
	pathName := GetPathName()
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

func GetChildren(c Child) []InfoDisplayed {
	var result []InfoDisplayed
	if c.Type == "folder" {
		for _, v := range c.Children {
			result = append(result, GetChildren(v)...)
		}
	} else {
		var pair InfoDisplayed
		pair.Name, pair.Url = c.Name, c.Url
		result = append(result, pair)
	}
	return result
}

func FilterByString(pairs []InfoDisplayed, search string) []InfoDisplayed {
	var result []InfoDisplayed
	for _, v := range pairs {
		isInName := strings.Contains(v.Name, search)
		isInUrl := strings.Contains(v.Url, search)
		if isInName || isInUrl {
			result = append(result, v)
		}
	}
	return result
}

func GetPathName() string {
	data, err := os.ReadFile("./settings.json")
	if err != nil {
		errColor := lipgloss.Color("#f52c43")
		fmt.Println(lipgloss.NewStyle().Foreground(errColor).Bold(true).SetString(err.Error()))
		fmt.Println(lipgloss.NewStyle().Foreground(errColor).Bold(true).SetString("Failed to read ./settings.json. Please check the setting file."))
		fmt.Println(lipgloss.NewStyle().Foreground(errColor).Bold(true).SetString("You can visit https://github.com/ebisuG/search-all-user-bookmark"))
		fmt.Print("Press 'Enter' to close...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}
	var settings SettingFile
	json.Unmarshal(data, &settings)
	var searchPathString strings.Builder
	searchPathString.WriteString("C:\\Users\\")
	searchPathString.WriteString(settings.UserName)
	searchPathString.WriteString("\\AppData\\Local\\Google\\Chrome\\User Data")
	return searchPathString.String()
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

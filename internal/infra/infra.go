package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ebisuG/search-all-user-bookmark/internal/config"
	"github.com/ebisuG/search-all-user-bookmark/internal/search"
	"golang.org/x/text/cases"
)

// This struct is for CLI settings
type ChromeLoader struct{}
type ChromeFinder struct{}

func (c ChromeLoader) Load(path string) (config.CliSetting, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return config.CliSetting{}, errors.New("failed to load config")
	}

	var cliSetting config.CliSetting
	if err := json.Unmarshal(data, &cliSetting); err != nil {
		fmt.Println(err)
		return config.CliSetting{}, errors.New("failed to parse json")
	}
	return cliSetting, nil
}

func (c ChromeFinder) Find(cliSetting config.CliSetting) (config.SearchPath, error) {
	base := "C:\\Users\\" + cliSetting.UserName + "\\AppData\\Local\\Google\\Chrome\\User Data"
	files, err := os.ReadDir(base)
	if err != nil {
		panic(err)
	}
	var bookmarksFilePath config.SearchPath
	bookmarksFilePath = append(bookmarksFilePath, base+"\\"+"Default"+"\\Bookmarks")
	r, _ := regexp.Compile("^Profile [0-9]*")

	for _, v := range files {
		match := r.MatchString(v.Name())
		if v.IsDir() && match {
			bookmarksFilePath = append(bookmarksFilePath, base+"\\"+v.Name()+"\\Bookmarks")
		}
	}

	return bookmarksFilePath, nil
}

type ChromeParentJson struct {
	Checksum     string      `json:"checksum"`
	Roots        ChromeRoots `json:"roots"`
	SyncMetadata string      `json:"sync_metadata"`
	Version      int         `json:"version"`
}

type ChromeRoots struct {
	BookmarkBar ChromeBookmarkBar `json:"bookmark_bar"`
}

type ChromeBookmarkBar struct {
	Children []ChromeChild `json:"children"`
	Other    ChromeChild   `json:"other"`
	Synced   ChromeChild   `json:"synced"`
}

type ChromeChild struct {
	Children     []ChromeChild  `json:"children"`
	DateAdded    string         `json:"date_added"`
	DateLastUsed string         `json:"date_last_used"`
	Guid         string         `json:"guid"`
	Id           string         `json:"id"`
	MetaInfo     ChromeMetaInfo `json:"meta_info"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	Url          string         `json:"url"`
}

type ChromeMetaInfo struct {
	PowerBookmarkMeta string `json:"power_bookmark_meta"`
}

type ChromeParser struct{}

func (c ChromeParser) Parse(path string) ([]search.Bookmark, error) {
	var bookmarks []search.Bookmark
	data, err := os.ReadFile(path)
	if err != nil {
		return bookmarks, errors.New("no file")
	}
	var chromeJson ChromeParentJson
	json.Unmarshal(data, &chromeJson)
	for i := 0; i < len(chromeJson.Roots.BookmarkBar.Children); i++ {
		bookmark := chromeJson.Roots.BookmarkBar.Children[i]
		bookmarks = append(bookmarks, GetChildren(bookmark)...)
	}
	return bookmarks, nil
}

func GetChildren(c ChromeChild) []search.Bookmark {
	folder := cases.Fold()
	var result []search.Bookmark
	if c.Type == "folder" {
		for _, v := range c.Children {
			result = append(result, GetChildren(v)...)
		}
	} else {
		var pair search.Bookmark
		pair.BookmarkTitle.Record.Raw = c.Name
		pair.BookmarkTitle.Record.Norm = folder.String(c.Name)
		pair.BookmarkUrl.Record.Raw = c.Url
		pair.BookmarkUrl.Record.Norm = folder.String(c.Url)
		result = append(result, pair)
	}
	return result
}

type CoreSearcher struct{}

func (c CoreSearcher) Search(bookmarks search.Bookmarks, keyword string) (search.Bookmarks, error) {
	var result search.Bookmarks
	folder := cases.Fold()
	searchWord := folder.String(keyword)
	for _, v := range bookmarks {
		isInName := strings.Contains(v.BookmarkTitle.Record.Norm, searchWord)
		isInUrl := strings.Contains(v.BookmarkUrl.Record.Norm, searchWord)
		if isInName || isInUrl {
			result = append(result, v)
		}
	}
	return result, nil
}

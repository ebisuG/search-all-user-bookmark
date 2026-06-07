package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
)

type Bookmarks []Bookmark

type Bookmark struct {
	BookmarkTitle BookmarkTitle
	BookmarkUrl   BookmarkUrl
}

type BookmarkTitle struct {
	Record Record
}

type BookmarkUrl struct {
	Record Record
}

type Record struct {
	Raw  string
	Norm string
}

type Searcher interface {
	Search(keyword string) (Bookmarks, error)
}

var _ Searcher = (*Bookmarks)(nil)

func (b *Bookmarks) Search(keyword string) (Bookmarks, error) {
	var result Bookmarks
	folder := cases.Fold()
	searchWord := folder.String(keyword)
	for _, v := range *b {
		isInName := strings.Contains(v.BookmarkTitle.Record.Norm, searchWord)
		isInUrl := strings.Contains(v.BookmarkUrl.Record.Norm, searchWord)
		if isInName || isInUrl {
			result = append(result, v)
		}
	}
	return result, nil
}

func NewChromeBookmarkSearcher(path []string) Searcher {
	var chromeBookmarks Bookmarks
	parser := NewChromeParser()
	for _, v := range path {
		bookmarks, err := parser.Parse(v)
		if err != nil {
			fmt.Println(err)
		}
		chromeBookmarks = append(chromeBookmarks, bookmarks...)
	}
	return &chromeBookmarks
}

type Parser interface {
	Parse(path string) (Bookmarks, error)
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

var _ Parser = (*ChromeParentJson)(nil)

func (c *ChromeParentJson) Parse(path string) (Bookmarks, error) {
	var bookmarks Bookmarks
	data, err := os.ReadFile(path)
	if err != nil {
		return bookmarks, errors.New("no file")
	}
	if err := json.Unmarshal(data, &c); err != nil {
		return bookmarks, errors.New("Invalid format")
	}
	for i := 0; i < len(c.Roots.BookmarkBar.Children); i++ {
		bookmark := c.Roots.BookmarkBar.Children[i]
		bookmarks = append(bookmarks, GetChildren(bookmark)...)
	}
	return bookmarks, nil
}

func GetChildren(c ChromeChild) Bookmarks {
	folder := cases.Fold()
	var result Bookmarks
	if c.Type == "folder" {
		for _, v := range c.Children {
			result = append(result, GetChildren(v)...)
		}
	} else {
		var pair Bookmark
		pair.BookmarkTitle.Record.Raw = c.Name
		pair.BookmarkTitle.Record.Norm = folder.String(c.Name)
		pair.BookmarkUrl.Record.Raw = c.Url
		pair.BookmarkUrl.Record.Norm = folder.String(c.Url)
		result = append(result, pair)
	}
	return result
}

func NewChromeParser() Parser {
	return &ChromeParentJson{}
}

// type Initializer interface {
// 	Initialize(path string) (Bookmarks, error)
// }

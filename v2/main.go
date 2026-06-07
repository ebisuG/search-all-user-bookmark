package main

import (
	"fmt"

	root "github.com/ebisuG/search-all-user-bookmark/v2/cmd"
	"github.com/ebisuG/search-all-user-bookmark/v2/internal/config"
	"github.com/ebisuG/search-all-user-bookmark/v2/internal/search"
)

func main() {
	loader := NewLoader()
	cfg := loader.Load()
	fmt.Println(cfg)

	chromeSearcher := NewChrome(cfg.Environments.BookmarkFilePaths)
	result, _ := chromeSearcher.Search("GitHub")
	fmt.Println(result)

	root.Execute()
}

func NewLoader() config.Loader {
	return &config.Config{}
}

func NewChrome(path []string) search.Searcher {
	return search.NewChromeBookmarkSearcher(path)
}

package cmd

import (
	"fmt"

	"github.com/ebisuG/search-all-user-bookmark/v2/internal/config"
	"github.com/ebisuG/search-all-user-bookmark/v2/internal/search"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "execute search keyword in all bookmark files",
	Run: func(cmd *cobra.Command, args []string) {
		loader := NewLoader()
		cfg := loader.Load()
		chromeSearcher := NewChrome(cfg.Environments.BookmarkFilePaths)
		result, _ := chromeSearcher.Search(args[0])
		fmt.Println(result)
	},
}

func NewLoader() config.Loader {
	return &config.Config{}
}

func NewChrome(path []string) search.Searcher {
	return search.NewChromeBookmarkSearcher(path)
}

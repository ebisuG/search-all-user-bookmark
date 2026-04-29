package config

import "github.com/ebisuG/search-all-user-bookmark/internal/search"

type Config struct {
	SearchPath SearchPath
	CliSetting CliSetting
}

type SearchPath []string

type CliSetting struct {
	UserName string `json:"username"`
}

type Loader interface {
	Load(path string) (CliSetting, error)
}

type Finder interface {
	Find(cli CliSetting) (SearchPath, error)
}

type Parser interface {
	Parse(path string) ([]search.Bookmark, error)
}

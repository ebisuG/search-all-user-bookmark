package config

import "github.com/ebisuG/search-all-user-bookmark/internal/search"

type Config struct {
	SearchPath []string
	CliSetting CliSetting
}

type CliSetting struct {
	UserName string `json:"username"`
}

type Loader interface {
	Load(path string) (Config, error)
}

type Finder interface {
	Find(conf Config) ([]string, error)
}

type Parser interface {
	Parse(path string) ([]search.Bookmark, error)
}

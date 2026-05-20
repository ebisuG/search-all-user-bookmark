package config

import "errors"

type Config struct {
	SearchPath SearchPath
	CliSetting CliSetting
}

type SearchPath []string

type CliSetting struct {
	UserName string `json:"USERNAME"`
	LogLevel string `json:"LOG_LEVEL"`
}

type Loader interface {
	Load(path string) (CliSetting, error)
}

type Finder interface {
	Find(cli CliSetting) (SearchPath, error)
}

type NotFoundError struct {
	FilePath string
}

func (e *NotFoundError) Error() string { return "No such file : " + e.FilePath }

var ErrInvalidFormat = errors.New("invalid format")

type NotFoundBookmarkFileError struct {
	Path string
}

func (e *NotFoundBookmarkFileError) Error() string {
	return "not found : " + e.Path
}

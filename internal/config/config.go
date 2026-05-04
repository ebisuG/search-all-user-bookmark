package config

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

type NotFoundError struct {
	FilePath string
}

func (e *NotFoundError) Error() string { return "No such file : " + e.FilePath }

type InvalidFormatError struct{}

func (e *InvalidFormatError) Error() string { return "Invalid format" }

type NotFoundBookmarkFileError struct {
	Path string
}

func (e *NotFoundBookmarkFileError) Error() string {
	return "not found : " + e.Path
}

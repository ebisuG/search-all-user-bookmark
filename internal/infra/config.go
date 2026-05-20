package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/ebisuG/search-all-user-bookmark/internal/config"
)

type ChromeLoaderImpl struct{}
type ChromeFinderImpl struct{}

var _ config.Loader = (*ChromeLoaderImpl)(nil)
var _ config.Finder = (*ChromeFinderImpl)(nil)

func (c *ChromeLoaderImpl) Load(path string) (config.CliSetting, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		//TODO : if debug mode on, display detailed data
		return config.CliSetting{}, fmt.Errorf("%w", config.NotFoundError{FilePath: path})
	}

	var cliSetting config.CliSetting
	if err := json.Unmarshal(data, &cliSetting); err != nil {
		return config.CliSetting{}, fmt.Errorf("%w", config.ErrInvalidFormat)
	}
	return cliSetting, nil
}

func NewChromeLoader() config.Loader {
	return &ChromeLoaderImpl{}
}

func (c *ChromeFinderImpl) Find(cliSetting config.CliSetting) (config.SearchPath, error) {
	base := "C:\\Users\\" + cliSetting.UserName + "\\AppData\\Local\\Google\\Chrome\\User Data"
	files, err := os.ReadDir(base)
	if err != nil {
		return config.SearchPath{}, errors.New("no such directory : " + base)
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
	if len(bookmarksFilePath) == 0 {
		return config.SearchPath{}, fmt.Errorf("%w", config.NotFoundBookmarkFileError{Path: base})
	}
	return bookmarksFilePath, nil
}

func NewChromeFinder() config.Finder {
	return &ChromeFinderImpl{}
}

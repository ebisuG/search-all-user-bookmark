package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ebisuG/search-all-user-bookmark/internal/config"
)

// This struct is for CLI settings
type BrowserLoader struct{}
type ChromeLoader struct{}
type ChromeFinder struct{}

func (c ChromeLoader) Load(path string) (config.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return config.Config{}, errors.New("failed to load config")
	}

	var conf config.Config
	if err := json.Unmarshal(data, &conf.CliSetting); err != nil {
		fmt.Println(err)
		return config.Config{}, errors.New("failed to parse json")
	}
	return conf, nil
}

//Define multiple struct for each browsers,
//Chrome, Edge, ...
//Then, define Find to each struct.

func (c ChromeFinder) Find(conf config.Config) ([]string, error) {
	base := "C:\\Users\\" + conf.CliSetting.UserName + "\\AppData\\Local\\Google\\Chrome\\User Data"
	return []string{base}, nil
}

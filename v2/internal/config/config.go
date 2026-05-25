package config

import (
	"encoding/json"
	"os"
	"regexp"
)

var _ Loader = (*Config)(nil)

type Config struct {
	CliSetting   CliSetting
	Environments Environments
}

type CliSetting struct {
	UserName string `json:"USERNAME"`
	LogLevel string `json:"LOG_LEVEL"`
}

type Environments struct {
	BookmarkFilePaths []string
}

func NewConfig(cli CliSetting, env Environments) *Config {
	return &Config{CliSetting: cli, Environments: env}
}

func (c *Config) LoadCliSettings(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, &c); err != nil {
		return
	}
}

func (c *Config) LoadEnvironment() {
	base := "C:\\Users\\" + c.CliSetting.UserName + "\\AppData\\Local\\Google\\Chrome\\User Data"
	files, err := os.ReadDir(base)
	if err != nil {
		return
	}
	c.Environments.BookmarkFilePaths = append(c.Environments.BookmarkFilePaths, base+"\\"+"Default"+"\\Bookmarks")
	r, _ := regexp.Compile("^Profile [0-9]*")

	for _, v := range files {
		match := r.MatchString(v.Name())
		if v.IsDir() && match {
			c.Environments.BookmarkFilePaths = append(c.Environments.BookmarkFilePaths, base+"\\"+v.Name()+"\\Bookmarks")
		}
	}
	if len(c.Environments.BookmarkFilePaths) == 0 {
		return
	}
}

func (c *Config) Load() Config {
	c.LoadCliSettings("./settings.json")
	c.LoadEnvironment()
	return *c //pointer dereference
}

type Loader interface {
	Load() Config
}

//Next TODO:
//main.go
// loader := NewLoader()
//cfg := loader.Load()
//domain.go
// domainLogic(cfg config.Config)
//move Loader interface to domain for "dependency inversion".

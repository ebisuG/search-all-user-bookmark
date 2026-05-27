package main

import (
	"fmt"

	root "github.com/ebisuG/search-all-user-bookmark/v2/cmd"
	"github.com/ebisuG/search-all-user-bookmark/v2/internal/config"
)

func main() {
	loader := NewLoader()
	cfg := loader.Load()
	fmt.Println(cfg)
	root.Execute()
}

func NewLoader() config.Loader {
	return &config.Config{}
}

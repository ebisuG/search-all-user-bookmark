package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	// "github.com/ebisuG/search-all-user-bookmark/internal/bubbleTea"

	"github.com/ebisuG/search-all-user-bookmark/internal/cli"
	"github.com/ebisuG/search-all-user-bookmark/internal/config"
	"github.com/muesli/termenv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	p := tea.NewProgram(cli.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	searchString textinput.Model
	config       config.Config
}

type hit struct {
	name          string
	url           string
	bookmarkTitle bookmarkTitle
	bookmarkUrl   bookmarkUrl
}

type result []hit

type bookmarkTitle struct {
	record record
}

type bookmarkUrl struct {
	record record
}

type record struct {
	raw  string
	norm string
}

type ConfigLoader struct{}

func (c *ConfigLoader) Find(conf config.Config) ([]string, error) {
	base := "C:\\Users\\" + conf.CliSetting.UserName + "\\AppData\\Local\\Google\\Chrome\\User Data"
	return []string{base}, nil
}

func (c *ConfigLoader) LoadConfig(path string) (config.Config, error) {
	var cmdConfig config.Config
	data, err := os.ReadFile(path)
	if err != nil {
		return config.Config{}, errors.New("failed to read file")
	}
	if err := json.Unmarshal(data, cmdConfig.CliSetting); err != nil {
		return config.Config{}, errors.New("failed to parse json")
	}
	paths, err := c.Find(cmdConfig)
	if err != nil {
		return config.Config{}, errors.New("failed to find target path")
	}
	cmdConfig.SearchPath = paths
	return cmdConfig, nil
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Search keyword"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#a871f0")).Bold(true).SetString("Press Ctrl + c when you quit"))
	fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#a871f0")).Bold(true).SetString("Press Ctrl to jummp to bookmark"))
	fmt.Println("Reading all bookmark files...")

	var modelConf ConfigLoader
	var conf config.Config

	conf, err := modelConf.LoadConfig("./settings.json")
	if err != nil {
		return model{}

	}

	return model{searchString: ti, config: conf}

}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			fmt.Println("Bye bye!")
			return m, tea.Quit

		case "enter", " ":
			//TODO: call search logic, with only interface
			// var searchWord []string
			// for _, v := range m.searchString.Value() {
			// 	searchWord = append(searchWord, string(v))
			// }
			// var display []search.InfoDisplayed
			// display = search.FilterByString(m.allUrl, strings.Join(searchWord, ""))
			// FormatDisplay(display)
			// m.searchString.Reset()
			// fmt.Println("")
		}
	}

	m.searchString, cmd = m.searchString.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := "Below string is being searched... \n\n"
	s += fmt.Sprintf("%s\n", m.searchString.View())

	s += "\nPress q to quit.\n"

	return fmt.Sprintf("%s\n", m.searchString.View())
}

func (r result) FormatDisplay() {
	nameColor := lipgloss.Color("#F77F0F")

	for _, v := range r {
		hypelink := termenv.Hyperlink(v.bookmarkUrl.record.raw, v.bookmarkTitle.record.raw)
		fmt.Println(" ", lipgloss.NewStyle().Foreground(nameColor).Bold(false).Render(hypelink))
	}
}

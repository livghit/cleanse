package main

import (
	"fmt"
	"os"

	"github.com/livghit/cleanse/cmd"
	"gopkg.in/yaml.v3"
)

type conf struct {
	SearchPaths   []string `yaml:"search_paths"`
	DebugKeywords []string `yaml:"debug_keywords"`
	Preset        string   `yaml:"preset"`
}

func (c *conf) loadConf() *conf {
	yamlFile, err := os.ReadFile("./cleanse.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}

	return c
}

var c conf

func main() {
	c.loadConf()
	if c.Preset == "" {
		fmt.Printf("No preset found trying to use custom config cleanse.yml file")
		searchDebugingStatementsViaCustomConfig()

	} else {
		// this will return the preset config need also to parse it
		loadPreset()
		searchDebugingStatementsViaPreset()
	}
}

func searchDebugingStatementsViaCustomConfig() {
	var keywords string
	var search string
	if c.Preset != "" {
		for _, k := range c.DebugKeywords {
			keywords += k
		}
		for _, s := range c.SearchPaths {
			search += s
		}
	}

	command := cmd.Command{
		Name: "grep -rEoHn",
		Args: []string{keywords, search},
	}
	output, err := command.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(output)
}

func searchDebugingStatementsViaPreset() {
	fmt.Println("Preset searching")
}

func loadPreset() {
	// Search for the presets atm in a folder later in git repo and we search for it there
	preset, err := os.ReadFile(fmt.Sprintf("./presets/%s.yml", c.Preset))
	if err != nil {
		fmt.Printf("Preset %s not found.", c.Preset)
	}

	fmt.Println("Found preset:")
	fmt.Println(preset)

}

package main

import (
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type conf struct {
	SearchPaths   []string `yaml:"search_paths"`
	DebugKeywords []string `yaml:"debug_keywords"`
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
	searchDebugingStatements()
}

// ATM just playing around not a implemented version
// Figuring out what I want to do
func searchDebugingStatements() {
	var keywords string
	var search string
	for _, k := range c.DebugKeywords {
		keywords += k
	}
	for _, s := range c.SearchPaths {
		search += s
	}
	cmd := exec.Command("sh", "-c", fmt.Sprintf("grep -rEoHn '%s' %s", keywords, search))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
}

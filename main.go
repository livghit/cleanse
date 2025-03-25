package main

import (
	"fmt"
	"os"

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

func searchDebugingStatements() {
	for _, path := range c.SearchPaths {
		fmt.Println(path)
	}
}

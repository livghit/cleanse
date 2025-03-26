package cmd

import (
	"os/exec"
	"strings"
)

type Cmd interface {
	Run() (string, error)
}

type Command struct {
	Name string
	Args []string
}

func (g *Command) Run() (string, error) {
	cmd := exec.Command(g.Name, g.Args...)
	output, err := cmd.CombinedOutput()

	return strings.TrimSpace(string(output)), err
}

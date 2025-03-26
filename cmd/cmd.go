package cmd

import "os/exec"

type Cmd interface {
	Command(name string, arg ...string) *exec.Cmd
}

package hacmd

import "os"

type Commander struct {
	Path string
}

func NewCommander() *Commander {
	return &Commander{}
}

func (cmd *Commander) SetPath(path string) *Commander {
	cmd.Path = path
	return cmd
}

func (cmd *Commander) Execute(name string, args ...string) {
	if cmd.Path != "" {
		os.Chdir(cmd.Path)
	}
	Execute(name, args...)
}

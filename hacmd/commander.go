package hacmd

import (
	"os"
)

type Commander struct {
	OriginPath string
	Path       string
}

func NewCommander() *Commander {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &Commander{
		OriginPath: currentPath,
	}
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

func (cmd *Commander) GoOrigin() error {
	return os.Chdir(cmd.OriginPath)
}

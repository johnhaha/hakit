package hacmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/johnhaha/hakit/hadata"
)

// execute with realtime output
func Execute(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Start()
	// fmt.Println("The command is running")
	if err != nil {
		fmt.Println(err)
	}
	b := hadata.NewStringBinder()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		b.BindWithNewLine(m)
		fmt.Println(m)
	}
	cmd.Wait()
	return b.Value()
}

// run cmd and return string output
func Run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.Output()
	ot := strings.TrimSuffix(string(stdout), "\n")
	return ot, err
}

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	ot, _ := reader.ReadString('\n')
	text := strings.Replace(ot, "\n", "", -1)
	return text
}

// run shell file
func Shell(path string, args ...string) {
	var eArgs []string
	eArgs = append(eArgs, path)
	eArgs = append(eArgs, args...)
	Execute("/bin/sh", eArgs...)
}

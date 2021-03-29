package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Jon1105/ProjectAutomation/common"
)

var cmds = [2]string{"create", "open"}
var target = filepath.Join(common.Documents, "Programming", "Scripts")

func main() {
	var command *exec.Cmd
	var src string

	for _, val := range cmds {
		command = exec.Command("go", "build", "./"+val)
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout

		fmt.Println(strings.Title(val), "script built successfully!")

		if err := command.Run(); err != nil {
			panic(err)
		}
		src = filepath.Join(common.CurrentPath, val+".exe")
		os.Rename(src, filepath.Join(target, val+".exe"))

		fmt.Println(strings.Title(val), "script moved successfully!")

	}
}

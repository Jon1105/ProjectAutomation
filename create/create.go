package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Jon1105/ProjectAutomation/common"
	"github.com/Jon1105/ProjectAutomation/git"
)

func main() {
	var nArgs int = len(os.Args)
	if nArgs < 2 || os.Args[1] == "help" {
		var helpFile string = filepath.Join(common.CurrentPath, "create", "help.txt")
		var data, err = ioutil.ReadFile(helpFile)
		common.Check(err)
		fmt.Print(string(data))
		return
	}

	// if nArgs >= 3 {...}

	// Establish boolean params
	var withGit bool = contains(os.Args, "-g")
	var private bool = !contains(os.Args, "-v")
	var local bool = contains(os.Args, "-l")
	var withREADME bool = contains(os.Args, "-r")

	// Get language object
	var lang, err1 = common.Classify(os.Args[1])
	common.Check(err1)

	var projectName string = os.Args[2]
	var projectPath string = filepath.Join(lang.Path, projectName)
	var templateFolder string = filepath.Join(common.CurrentPath, "create", "templates", lang.Name)

	var exists, err2 = common.Exists(filepath.Join(templateFolder, "template"))
	common.Check(err2)
	if exists {
		common.Check(copy(templateFolder, projectPath, []string{"template"}))
	}

	if lang.Name == "Arduino" {
		var newPath string = filepath.Join(projectPath, projectName+".ino")
		var err1 error = os.Rename(filepath.Join(projectPath, "template.ino"), newPath)

		common.Check(err1)
	} else if lang.Name == "Flutter" {
		var cmd = exec.Command("flutter", "create", projectPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		common.Check(cmd.Run())

		// Make necessary file changes
		common.Check(os.RemoveAll(filepath.Join(projectPath, "test")))

		// Alternative: For every file in template folder, replace or create the existing one
		var data, err1 = ioutil.ReadFile(filepath.Join(templateFolder, "lib", "main.dart"))
		common.Check(err1)
		common.Check(ioutil.WriteFile(filepath.Join(projectPath, "lib", "main.dart"), data, 0777))

	} else if lang.Name == "Node.js" {
		var cmd *exec.Cmd
		if contains(os.Args, "-y") {
			cmd = exec.Command("npm", "init", "-y")
		} else {
			cmd = exec.Command("npm", "init")
		}
		cmd.Dir = projectPath
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		common.Check(cmd.Run())
	} else if lang.Name == "Rust" {
		var cmd = exec.Command("cargo", "new", projectPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		common.Check(cmd.Run())

		// delete projectPath/.git

	} else if lang.Name == "Go" {
		var cmd = exec.Command("go", "mod", "init", projectName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = projectPath
		common.Check(cmd.Run())
	}

	if withREADME {
		var contents string = fmt.Sprintf("### %v\n---\n", projectName)
		var readmePath = filepath.Join(projectPath, "README.md")
		var err1 = ioutil.WriteFile(readmePath, []byte(contents), 0777)
		common.Check(err1)
	}

	var err3 error
	if withGit && local {
		err3 = git.Git(projectPath)
	} else if withGit && !local {
		err3 = git.Github(projectPath, private)
	}
	common.Check(err3)
	common.OpenWithCode(projectPath)

}

func contains(array []string, obj string) bool {
	for _, val := range array {
		if val == obj {
			return true
		}
	}
	return false
}

func copy(source string, destination string, omit []string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath string = strings.Replace(path, source, "", 1)
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		}
		for _, v := range omit {
			if v == info.Name() {
				return nil
			}
		}
		var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
		if err1 != nil {
			return err1
		}
		return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)

	})
}

package main

import (
	"ProjectAutomation/common"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	var nArgs int = len(os.Args)
	if nArgs < 2 || os.Args[1] == "help" {
		var helpFile string = filepath.Join(common.CurrentPath, "open", "help.txt")
		var data, err = ioutil.ReadFile(helpFile)
		common.Check(err)
		fmt.Print(string(data))
		return
	} else if nArgs == 2 {
		// Get language object
		var lang, err1 = common.Classify(os.Args[1])
		common.Check(err1)

		// Print all Projects
		var count int
		var files, err2 = ioutil.ReadDir(lang.Path)
		common.Check(err2)
		var projects []os.FileInfo
		for _, file := range files {
			if file.IsDir() {
				projects = append(projects, file)
				count++
				fmt.Printf("%v: %v\n", count, file.Name())
			}
		}

		var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
		for true {
			fmt.Print("Pick a project: ")
			scanner.Scan()
			var input string = scanner.Text()
			if num, err := strconv.Atoi(input); err == nil && num <= len(projects) {
				common.OpenWithCode(filepath.Join(lang.Path, projects[num-1].Name()))
				return
			} else if input == "q" {
				return
			} else {
				path := filepath.Join(lang.Path, input)
				var exists, err3 = common.Exists(path)
				common.Check(err3)
				if exists {
					common.OpenWithCode(path)
					return
				} else if !exists {
					fmt.Println("Invalid Entry")
				}
			}
		}
	} else if nArgs == 3 {
		var lang, err1 = common.Classify(os.Args[1])
		common.Check(err1)
		var path string = filepath.Join(lang.Path, os.Args[2])
		var exists, err2 = common.Exists(path)
		common.Check(err2)
		if exists {
			common.OpenWithCode(path)
			return
		} else if !exists {
			fmt.Println(fmt.Errorf("%q is not a valid path on this machine", path))
			return
		}
	} else {
		fmt.Println(fmt.Errorf("Excpected 3 or less arguments, got %d", nArgs))
		return 
	}

}

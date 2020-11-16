package common

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var home, err = os.UserHomeDir()

// Documents path to documents folder where projects lie
var Documents string = filepath.Join(home, "OneDrive", "Documents")

// GoPath Path to Go projects to find current directory
var GoPath string = filepath.Join(Documents, "Programming", "Go", "src")

// CurrentPath Path representing location of this files parent's parent's
var CurrentPath string = filepath.Join(GoPath, "ProjectAutomation")

// Language struct to hold all variable info regarding each language with projects folder
type Language struct {
	Name, Path string
}

var allLanguages = map[string]Language{
	"Python":  Language{"Python", w("Programming/Python/Projects")},
	"Go":      Language{"Go", GoPath},
	"Node.js": Language{"Node.js", w("Programming/Node.js/Projects")},
	"C++":     Language{"C++", w("Programming/C++/Projects")},
	"Flutter": Language{"Flutter", w("Programming/Flutter/Projects")},
	"Arduino": Language{"Arduino", w("Electronics/Arduino/Sketches")},
	"Rust": Language{"Rust", w("Programming/Rust/Projects")},
}

func w(path string) string {
	return filepath.Join(Documents, path)
}

// Classify used to determine the language from human inputted strings
func Classify(strLang string) (Language, error) {
	var key string
	var err error
	if strings.Contains(strings.ToLower(strLang), "py") {
		key = "Python"
	} else if strings.Contains(strings.ToLower(strLang), "c") {
		key = "C++"
	} else if strings.Contains(strings.ToLower(strLang), "node") {
		key = "Node.js"
	} else if strings.Contains(strings.ToLower(strLang), "flutter") {
		key = "Flutter"
	} else if strings.Contains(strings.ToLower(strLang), "go") {
		key = "Go"
	} else if strings.Contains(strings.ToLower(strLang), "ino") {
		key = "Arduino"
	} else if strings.Contains(strings.ToLower(strLang), "rust") {
		key = "Rust"
	} else {
		err = fmt.Errorf("common: string %q does not represent a valid language", strLang)
	}
	return allLanguages[key], err
}

// Exists function to check whether path exists on the machine
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// OpenWithCode wrapper to open a folder with vscode
func OpenWithCode(path string) error {
	var cmd *exec.Cmd = exec.Command("code", path)
	return cmd.Run()
}

// Check wrapper for checking for errors and reporting them immediately
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

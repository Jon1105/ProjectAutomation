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
var GoPath string = filepath.Join(Documents, "Programming", "Go", "src", "github.com", "Jon1105")

// CurrentPath Path representing location of this files parent's parent's
var CurrentPath string = filepath.Join(GoPath, "ProjectAutomation")

// Language struct to hold all variable info regarding each language with projects folder
type Language struct {
	Name, Path string
}

func docs(path string) string {
	return filepath.Join(Documents, path)
}

// Classify used to determine the language from human inputted strings
func Classify(strLang string) (Language, error) {
	var lang Language
	var err error
	var low string = strings.ToLower(strLang)
	if low == "py" || low == "python" {
		lang = Language{"Python", docs("Programming/Python/Projects")}

	} else if low == "cpp" || low == "c++" {
		lang = Language{"C++", docs("Programming/C++/")}

	} else if low == "node" || low == "js" || low == "ts" {
		lang = Language{"Node.js", docs("Programming/Node.js/")}

	} else if low == "flutter" || low == "dart" {
		lang = Language{"Flutter", docs("Programming/Flutter/")}

	} else if low == "go" {
		lang = Language{"Go", GoPath}

	} else if low == "ino" || low == "arduino" {
		lang = Language{"Arduino", docs("Electronics/Arduino/Sketches")}

	} else if low == "rust" {
		lang = Language{"Rust", docs("Programming/Rust/")}

	} else if low == "workspace" {
		lang = Language{"Workspace", filepath.Join(home, ".vscode", "Workspaces")}

	} else {
		err = fmt.Errorf("common: string %q does not represent a valid language", strLang)

	}
	return lang, err
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

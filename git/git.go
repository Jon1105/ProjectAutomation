package git

import (
	"ProjectAutomation/common"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const githubUsername string = "Jon1105"

func githubRepo(repoName string, private bool) (string, int, error) {
	var key, err1 = ioutil.ReadFile(filepath.Join(common.CurrentPath, "git", "github.key"))

	if err1 != nil {
		return "", 0, err1
	}
	var data = map[string]interface{}{
		"name":    repoName,
		"private": private,
	}
	var raw, err2 = json.Marshal(data)
	if err2 != nil {
		return "", 0, err2
	}

	const url string = "https://api.github.com/user/repos"
	var request, err3 = http.NewRequest("POST", url, bytes.NewBuffer(raw))
	if err3 != nil {
		return "", 0, err3
	}
	request.SetBasicAuth(githubUsername, string(key))
	fmt.Println("Request Headers:", request.Header)
	var client *http.Client = &http.Client{}
	var response, err4 = client.Do(request)
	if err4 != nil {
		return "", 0, err4
	}

	defer response.Body.Close()
	fmt.Println("")
	fmt.Println("Response Status:", response.Status)
	fmt.Println("")
	fmt.Println("Response Header:", response.Header)
	fmt.Println("")
	var body, err5 = ioutil.ReadAll(response.Body)
	if err5 != nil {
		return "", response.StatusCode, err5
	}
	var bodyMap map[string]interface{}

	if err6 := json.Unmarshal(body, &bodyMap); err6 != nil {
		return "", response.StatusCode, err6
	}

	return fmt.Sprintf("%v", bodyMap["git_url"]), response.StatusCode, nil
}

// Git Used to create local git repositories
func Git(projectPath string) error {
	var cmd *exec.Cmd = exec.Command("git", "init", projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Github Used to create remote github repositories and synchronise with local
func Github(projectPath string, private bool) error {
	if err1 := Git(projectPath); err1 != nil {
		return err1
	}
	var url, statusCode, err2 = githubRepo(filepath.Base(projectPath), private)
	if err2 != nil {
		return err2
	}
	if (statusCode >= 200 && statusCode <= 299) && url != "" {
		var cmd *exec.Cmd = exec.Command("git", "remote", "add", "origin", url)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	return fmt.Errorf("%d: %s", statusCode, http.StatusText(statusCode))
}

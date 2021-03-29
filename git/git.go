package git

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Jon1105/ProjectAutomation/common"
	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
)

func getKey() (string, error) {
	var key, err1 = ioutil.ReadFile(filepath.Join(common.CurrentPath, "git", "github.key"))
	if err1 != nil {
		return "", err1
	} else {
		return string(key), nil
	}
}

func createRepo(repoName string, private bool) (*github.Repository, error) {
	key, err1 := getKey()
	if err1 != nil {
		return nil, err1
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: key},
	)

	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(private),
	}
	resRepo, _, err2 := client.Repositories.Create(context.Background(), "", repo)
	return resRepo, err2
}

func Git(projectPath string) error {
	var cmd *exec.Cmd = exec.Command("git", "init", projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Github(projectPath string, private bool) error {
	if err1 := Git(projectPath); err1 != nil {
		return err1
	}
	repo, err2 := createRepo(filepath.Base(projectPath), private)
	if err2 != nil {
		return err2
	}

	var cmd *exec.Cmd = exec.Command("git", "remote", "add", "origin", *repo.GitURL)
	cmd.Dir = projectPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Repository struct {
	dir           path
	status        string
	currentBranch branch
	projectType   string
	branches      []branch
	children      []Repository
}

func (repo *Repository) Dir() string {
	return repo.dir.String()
}

func (repo *Repository) runGitCommand(param ...string) string {
	curDir, _ := os.Getwd()
	os.Chdir(repo.dir.String())
	defer os.Chdir(curDir)

	cmd := exec.Command("git", param...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	return out.String()
}

func (repo *Repository) Status() string {
	repo.status = repo.runGitCommand("status", "--porcelain")
	fmt.Println(repo.status)
	return repo.status
}

func (repo *Repository) Branches() []string {
	var branches []string
	for _, value := range repo.getBranches() {
		branches = append(branches, value.name)
	}

	return branches
}

func (repo *Repository) getBranches() []branch {
	branches := strings.Split(repo.runGitCommand("branch"), "\n")
	var result []branch
	for _, value := range branches {
		isMaster := false
		if strings.Contains(value, "*") {
			isMaster = true
		}
		tempName := strings.Trim(value, " *")
		if len(tempName) > 0 {
			result = append(result, branch{tempName})
		}
		if isMaster {
			repo.currentBranch = branch{tempName}
		}
	}
	repo.branches = result

	return repo.branches
}

func (repo *Repository) setProjectType() {
	if _, err := os.Stat(repo.dir.add("/typo3conf").String()); err == nil {
		repo.projectType = "TYPO3"
	}
}

func (repo *Repository) findChildren() (result []Repository) {
	if repo.projectType != "TYPO3" {
		tempDir, _ := os.Getwd()
		os.Chdir(repo.dir.add("/typo3conf/ext/").String())
		defer os.Chdir(tempDir)

	}
	return
}

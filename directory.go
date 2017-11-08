package git

import "os"

type Directory struct {
	absolutePath path
}

func (dir *Directory) WorkingDir() path {
	var tmp string
	tmp, _ = os.Getwd()
	dir.absolutePath = path{tmp}

	return dir.absolutePath
}

func (dir *Directory) isGitDirPath() bool {
	gitDir := dir.absolutePath.add("/.git")
	if _, err := os.Stat(gitDir.String()); os.IsNotExist(err) {
		return false
	}

	return true
}

func (dir *Directory) createRepository() (repo *Repository) {
	repo = new(Repository)
	repo.dir = dir.absolutePath
	repo.setProjectType()

	return
}

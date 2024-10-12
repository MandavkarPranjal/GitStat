package git

import (
	"gopkg.in/src-d/go-git.v4"
	// "log"
)

// Repository wraps the go-git repository
type Repository struct {
	*git.Repository
}

// OpenRepository opens a Git repository at the given path
func OpenRepository(path string) (*Repository, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	return &Repository{repo}, nil
}

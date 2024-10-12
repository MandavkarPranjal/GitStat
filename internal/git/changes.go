package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"log"
)

// Change represents a file change
type Change struct {
	FilePath string `json:"file_path"`
	Action   string `json:"action"` // Added, Modified, Deleted
}

// GetFileChanges retrieves file changes from the repository
func GetFileChanges(repo *Repository) ([]Change, error) {
	var changes []Change

	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	// Get the commit history
	commitIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, err
	}

	// Iterate over the commits
	err = commitIter.ForEach(func(c *object.Commit) error {
		if c.NumParents() == 0 {
			return nil // Skip initial commit
		}

		parent, err := c.Parent(0)
		if err != nil {
			return err
		}

		patch, err := parent.Patch(c)
		if err != nil {
			return err
		}

		for _, fileStat := range patch.Stats() {
			action := "Modified"
			if fileStat.Addition == 0 {
				action = "Deleted"
			} else if fileStat.Deletion == 0 {
				action = "Added"
			}

			changes = append(changes, Change{
				FilePath: fileStat.Name,
				Action:   action,
			})
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error iterating commits: %v", err)
	}

	return changes, nil
}

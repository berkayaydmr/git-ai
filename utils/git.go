package utils

import (
	"fmt"
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func PrintBranches(directory string) ([]string, error) {
	repo, err := git.PlainOpen(directory)
	if err != nil {
		return nil, err
	}

	refs, err := repo.Branches()
	if err != nil {
		return nil, err
	}

	var branches []string
	refs.ForEach(func(r *plumbing.Reference) error {
		branches = append(branches, r.String())
		fmt.Println(r.String())
		return nil
	})

	return branches, err
}

func BranchDiff(directory, branch, branch2 string) (string, error) {
	repo, err := git.PlainOpen(directory)
	if err != nil {
		return "", err
	}

	commit1, err := getBranchCommit(repo, branch)
	if err != nil {
		return "", err
	}

	commit2, err := getBranchCommit(repo, branch2)
	if err != nil {
		return "", err
	}

	tree1, err := commit1.Tree()
	if err != nil {
		return "", err
	}

	tree2, err := commit2.Tree()
	if err != nil {
		return "", err
	}

	changes, err := object.DiffTree(tree1, tree2)
	if err != nil {
		return "", err
	}

	patch, err := changes.Patch()
	if err != nil {
		return "", err
	}

	return patch.String(), err
}

func getBranchCommit(repo *git.Repository, branchName string) (*object.Commit, error) {
	ref, err := repo.Reference(plumbing.NewBranchReferenceName(branchName), true)
	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	return commit, nil
}

func GetGitRepo(dir string) (*config.Config, error) {
	gitDir := filepath.Join(dir, ".git")
	info, err := os.Stat(gitDir)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, nil
	}

	repo, err := git.PlainOpen(dir)
	if err != nil {
		return nil, err
	}

	return repo.Config()
}

func CheckBranchExist(directory string, branch string) bool {
	if branch == "" {
		return false
	}

	repo, err := git.PlainOpen(directory)
	if err != nil {
		fmt.Println("118", err)
		return false
	}

	branches, err := repo.Branches()
	if err != nil {
		fmt.Println("120", err)
		return false
	}

	var branchExist bool
	branches.ForEach(func(r *plumbing.Reference) error {
		if r.Name().Short() == branch {
			branchExist = true
		}
		return nil
	})

	return branchExist
}

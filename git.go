package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Git struct {
	repo *git.Repository
}

// cloneRepo pulls down the repo given at the repoURI from the flag arguments into the directory given
func cloneRepo() error {
	fmt.Println("start clone repo")

	cloneOptions := git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	}

	if checkoutFrom != nil {
		cloneOptions.ReferenceName = plumbing.ReferenceName(*checkoutFrom)
	}

	r, err := git.PlainClone(dirName, false, &cloneOptions)
	if err != nil {
		fmt.Printf("could not clone repo: %v\n", err)
		return err
	}

	fmt.Println("repo cloned")
	repo := Git{ repo: r}
	// checkout a new branch
	if err := repo.CreateNewBranch(); err != nil {
		log.Printf("could not create new branch: %v\n", err)
		return err
	}
	return nil
}

// createNewBranch checks out a branch with the given name
func (g Git) CreateNewBranch() error {
	log.Printf("checking out new branch: %v\n", newBranchName)
	// create a new work tree based on the named directory
	w, err := g.repo.Worktree()
	if err != nil {
		log.Printf("could not create worktree: %v\n", err)
		return err
	}

	// and checkout a new branch based on the given name
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(newBranchName),
		Create: true,
	})
	if err != nil {
		log.Printf("could not checkout branch %v\n", err)
		return err
	}
	log.Println("new branch created")
	return nil
}

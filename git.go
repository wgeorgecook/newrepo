package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

type Git struct {
	repo *git.Repository
}

// cloneRepo pulls down the repo given at the repoURI from the flag arguments into the directory given
func cloneRepo() error {
	log.Println("start clone repo")

	cloneOptions := git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	}

	var r *git.Repository
	var err error
	r, err = git.PlainClone(dirName, false, &cloneOptions)
	if err != nil {
		log.Printf("could not clone repo: %v\n", err)
		return err
	}

	if checkoutFrom != nil {
		log.Println("fetching branches on remote")
		w, _ := r.Worktree()
		err := r.Fetch(&git.FetchOptions{
			RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
		})
		if err != nil {
			fmt.Printf("could not fetch existing branches: %v\n", err)
			return err
		}
		log.Printf("checking out given branch %v\n", *checkoutFrom)
		err = w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", *checkoutFrom)),
			Force:  true,
		})
		if err != nil {
			fmt.Printf("could not checkout given branch: %v", err)
			return err
		}
	}

	log.Println("repo cloned")
	repo := Git{repo: r}
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
		Force:  true,
	})
	if err != nil {
		log.Printf("could not checkout branch: %v\n", err)
		return err
	}
	log.Println("new branch created")
	return nil
}

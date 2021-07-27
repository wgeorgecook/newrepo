package main

import (
	"flag"
	"log"
	"path/filepath"
)

var (
	repoURL       string
	checkoutFrom  *string
	newDirPath    string
	newBranchName string
	dirName       string
)

func init() {
	log.Println("start create new repo")

	// parse the incoming flags
	flag.StringVar(&repoURL, "r", "~/go/src", "URI for the main repo branch you want to start from")
	flag.StringVar(&newDirPath, "p", "new-dir", "Abosolute directory to clone repo from the repo URL into; creates a new directory with the given branch name")
	flag.StringVar(&newBranchName, "n", "newBranch", "Name to give new branch after cloning")
	checkoutFrom = flag.String("c", "", "Branch to checkout from after cloning before creating a new branch")
	flag.Parse()

	if *checkoutFrom == "" {
		checkoutFrom = nil
	}
	// set the new directory name
	dirName = filepath.Join(newDirPath, newBranchName)
}
func main() {
	log.Println("start new repo")
	// create a new directory and make that our working directory
	if err := createRepoDir(); err != nil {
		cleanup()
		panic("could not create repository directory")
	}

	// clone the main branch from the given URL and then checkout a new branch
	if err := cloneRepo(); err != nil {
		cleanup()
		panic("could not setup repository")
	}
	log.Println("complete")
}

package main

import (
	"log"
	"os"
)

// createRepoDir creates a new directory at the given path for the given name
func createRepoDir() error {
	// create the directory
	log.Printf("creating directory: %s\n", dirName)
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Printf("could not create directory: %v\n", err)
		return err
	}
	log.Printf("created directory: %s\n", dirName)

	// then navigate into it
	changeDirectory()
	return nil
}

// changeDirectory sets the current working directory to the new directory created
func changeDirectory() {
	log.Println("setting working directory")
	os.Chdir(dirName)
	log.Println("working directory set")
}

// cleanup removes created directory on failure
func cleanup() {
	log.Printf("start cleanup of %s\n", dirName)
	defer log.Println("end cleanup")
	if err := os.RemoveAll(dirName); err != nil {
		log.Printf("cleanup failed: %v\n", err)
		return
	}
}

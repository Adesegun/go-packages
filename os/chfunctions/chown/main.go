package main

import (
	"log"
	"os"
)

func main() {
	// Create a new file that is owned by the current user
	os.Create("file.txt")

	// chown the file file.txt
	//
	// Chown changes the numeric uid and gid of the named file. If the file
	// is a symbolic link, it changes the uid and gid of the link's target.
	// If there is an error, it will be of type *PathError.
	//
	// In this case chown won't work because there is no other user to
	// change the permissions to
	err := os.Chown("file.txt", 1, 1)

	// Log any errors if any occured
	if err != nil {
		log.Fatalln(err)
	}
}

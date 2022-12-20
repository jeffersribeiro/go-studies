package files_dir

import (
	"log"
	"os"
)

func MoveFile() {
	oldLocation := "packages/files_dir/files/renamed.txt"
	newLocation := "packages/files_dir/directories/moved.txt"
	err := os.Rename(oldLocation, newLocation)

	if err != nil {
		log.Fatal(err)
	}
}

package files_dir

import (
	"log"
	"os"
)

func RenameFile() {
	oldName := "packages/files_dir/files/empty.txt"
	newName := "packages/files_dir/files/renamed.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

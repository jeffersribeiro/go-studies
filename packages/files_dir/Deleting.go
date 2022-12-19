package files_dir

import (
	"log"
	"os"
)

func Deleting() {
	err := os.Remove("packages/files_dir/files/copied.txt")
	if err != nil {
		log.Fatal(err)
	}
}

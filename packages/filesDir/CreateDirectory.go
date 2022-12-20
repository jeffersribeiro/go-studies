package files_dir

import (
	"log"
	"os"
)

func CreateDirectory() {
	_, err := os.Stat("packages/files_dir/directories")

	if os.IsNotExist(err) {
		errDir := os.Mkdir("packages/files_dir/directories", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

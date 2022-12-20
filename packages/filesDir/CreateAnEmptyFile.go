/*
Files and directories with examples


*/

package files_dir

import (
	"log"
	"os"
)

func CreateEmptyFile() {
	emptyFile, err := os.Create("packages/files_dir/files/empty.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(emptyFile)
	emptyFile.Close()
}

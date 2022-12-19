package files_dir

import (
	"io"
	"log"
	"os"
)

func CopyFile() {
	sourceFile, err := os.Open("packages/files_dir/directories/moved.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("packages/files_dir/files/copied.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Copied %d bytes.", bytesCopied)
}

/*
Output:

2022/12/19 15:13:35 Copied 0 bytes.
*/

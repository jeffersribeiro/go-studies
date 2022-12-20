package files_dir

import (
	"fmt"
	"log"
	"os"
)

func GettingMetadata() {
	fileStat, err := os.Stat("packages/files_dir/files/copied.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file Name: ", fileStat.Name())        // File Name: test.txt
	fmt.Println("file Size: ", fileStat.Size())        // Size: 100
	fmt.Println("Permissions: ", fileStat.Mode())      // Permissions: -rw-rw-rw-
	fmt.Println("file Modified: ", fileStat.ModTime()) // Last Modified: 2018-08-11 20:19:14.2671925 +0530 IST
	fmt.Println("Is Directory: ", fileStat.IsDir())    // Is Directory:  false
}

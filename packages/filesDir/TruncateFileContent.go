package files_dir

import (
	"log"
	"os"
)

/*
os.Trancate() function will recude the file content up to N bytes passed in second parameter. In below
exmaple if size of test.txt is more than 1Kb(100 byte) then it will truncate the remainins content
*/
func TruncateFileContent() {
	err := os.Truncate("test.txt", 100)

	if err != nil {
		log.Fatal(err)
	}
}

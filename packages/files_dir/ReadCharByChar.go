package files_dir

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadCharByChar() {
	filename := "packages/files_dir/files/copied.txt"
	fileBuffer, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(fileBuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))

	data.Split(bufio.ScanRunes)

	for data.Scan() {
		fmt.Println(data.Text())
	}
}

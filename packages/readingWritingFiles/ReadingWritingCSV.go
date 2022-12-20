package reading_files

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadingCSV() {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var textlines []string

	for scanner.Scan() {
		textlines = append(textlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range textlines {
		fmt.Println(eachline)
	}
}

func WritingCSV() {
	rows := [][]string{
		{"Name", "City", "Language"},
		{"Pinky", "London", "Python"},
		{"Nicky", "Paris", "Golang"},
		{"Micky", "Tokyo", "Php"},
	}

	csvfile, err := os.Create("test.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)

	for _, row := range rows {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()

	csvfile.Close()
}

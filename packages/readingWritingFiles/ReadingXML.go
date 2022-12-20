package reading_files

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func ReadingXML() {
	data, _ := ioutil.ReadFile("packages/reading_writing _file/notes.xml")

	note := &Notes{}

	_ = xml.Unmarshal([]byte(data), &note)

	fmt.Println(note.To)      //Tove
	fmt.Println(note.From)    //Jani
	fmt.Println(note.Heading) //Reminder
	fmt.Println(note.Body)    //Don't forget me this weekend!
}

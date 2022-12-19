package logging

import (
	"log"
	"net/smtp"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func RealWorldLoggingExample() {
	client, err := smtp.Dial("smtp.smail.com:25")

	if err != nil {
		log.Fatalln(err)
	}

	client.Data()
}

/*
Output:

TRACE: 2022/12/19 14:53:16.704868 /mnt/HDD/Projects/go-examples/packages/logging/RealWorldLoggingExample.go:11: init started
TRACE: 2022/12/19 14:55:27.998117 /mnt/HDD/Projects/go-examples/packages/logging/RealWorldLoggingExample.go:18: dial tcp 47.91.14.220:25: connect: connection timed out
exit status 1
*/

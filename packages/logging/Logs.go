package logging

import "log"

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.Println("Init started")
}

func Logs() {
	log.Println("Main Started")

	log.Fatalln("Fatal Message")

	log.Panicln("Panic Message")
}

/*
Output:

LOG: 2022/12/19 14:46:05.381170 /mnt/HDD/Projects/go-examples/packages/logging/Logs.go:9: Init started
LOG: 2022/12/19 14:46:05.381233 /mnt/HDD/Projects/go-examples/packages/logging/Logs.go:13: Main Started
LOG: 2022/12/19 14:46:05.381238 /mnt/HDD/Projects/go-examples/packages/logging/Logs.go:15: Fatal Message
exit status 1
*/

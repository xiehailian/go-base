package main

import (
	"log"
	"os"
)

func main() {
	logfile, err := os.OpenFile("my.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("fail to create log file!")
	}
	defer logfile.Close()

	l:=log.New(logfile, "", log.LstdFlags)
	l.Println("test")
	num:=5
	l.Println("test %d",num)
}

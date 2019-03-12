package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func Init()  {
	fmt.Println("haha")
}

type Option func()


func Dsn(dsn string) Option {
	return Option(func() {
		fmt.Println(dsn)
	})
}

func Args(first int, arg ...interface{})  {
	fmt.Println(first, arg)
}

func main() {
	start := time.Now()
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	time.Sleep(time.Second)
	fmt.Println(time.Since(start))
	fmt.Println(float64(time.Since(start)) * 1e-6)
}

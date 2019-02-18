package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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

  log.WithFields(log.Fields{
    "animal": "walrus",
  }).Info("A walrus appears")

  nums := []int64{1, 2, 3, 4}
  Args(1, nums)

  Args(1)

  Args(1, 2)
}

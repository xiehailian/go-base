package main

import (
	"fmt"
)

type ByteSize float64

const (
    _ = iota // 通过赋值给空白标识符来忽略值
    KB ByteSize = 1<<(10*iota)  // 二进制1左移10位，1000000000
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
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
	//start := time.Now()
	//log.WithFields(log.Fields{
	//	"animal": "walrus",
	//}).Info("A walrus appears")
	//
	//time.Sleep(time.Second)
	//fmt.Println(time.Since(start))
	//fmt.Println(float64(time.Since(start)) * 1e-6)

	type class struct {
		i int
	}

	 m :=  make(map[string]*class)
	 c := &class{9}
	 m["hehe"] = c
	 c.i = 10
	fmt.Println(m["hehe"].i)

	fmt.Println(KB)


}

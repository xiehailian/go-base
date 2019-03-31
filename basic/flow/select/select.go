package main

import (
	"fmt"
	"time"
)

/*
如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行
 */

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int

	select {     // select的case语句必须是一个channel操作
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	case <-time.After(time.Second * 3): //超时退出
		fmt.Println("request time out")
	//default:

	}
}
package main

import (
	"fmt"
	"time"
)

func main() {

	// 不使用带缓冲区的channel, 发送和接收的同步的, 发送以后如果不接受就channel就赌了，不能再发送到通道了
	//c := make(chan int)
	//go send(c)
	//go recv(c)
	//time.Sleep(3 * time.Second)
	//close(c)

	// 使用带缓冲区的channel，定了一个缓冲区容量，通道是异步的, 发送以后不接受可以一直发，知道channel满了
	d := make(chan int, 10)
	go send(d)
	//go recv(d)
	time.Sleep(3 * time.Second)
	close(d)  // 关闭以后可以接收数据，但是不能发送数据
	fmt.Println(<-d)
	//d <- 1    // close以后不能接收数据，否则报错
}

// 只能向chan里send数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {

		fmt.Println("send readey ", i)
		c <- i
		fmt.Println("send ", i)
	}
}

// 只能接收channel中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received ", i)
	}
}
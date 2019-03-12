package main

import (
	"fmt"
	"time"
)

func main() {
	message_chan := make(chan interface{}, 1)
	fmt.Println(cap(message_chan))
	go func() {
		for {
			println("start recv...")
			//message_chan <- nil
			println("finish recv...")
		}
		time.Sleep(time.Second * 3)

	}()

	//println("start send 10...")
	//message_chan <- 10
	//
	//println("start send 20...")
	//message_chan <- 20
	//
	//println("start send 30...")
	//message_chan <- 30
	//
	//println("finish send...")

	time.Sleep(time.Second * 3)
	<-message_chan
}

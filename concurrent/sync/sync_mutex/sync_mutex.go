package main

import (
	"fmt"
	"sync"
	"time"
)



func main() {
	wg := sync.WaitGroup{}

	var mutex sync.Mutex   // 互斥锁
	fmt.Println("Locking  (G0)")
	mutex.Lock()
	fmt.Println("locked (G0)")
	wg.Add(3)    // 需要等待的goroute数量

	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("locked (G%d)\n", i)

			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("unlocked (G%d)\n", i)
			wg.Done()    // 相当于 wg.Add(-1)
		}(i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock (G0)")
	mutex.Unlock()
	fmt.Println("unlocked (G0)")

	wg.Wait()     // 用来阻塞，当所有线程执行完毕，释放
}

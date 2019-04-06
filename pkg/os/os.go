package main


import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGUSR2)
	for {
		s := <-c
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
	}
}

func main() {
	// 1) os.StartProcess //
	/*********************/
	/* Linux: */
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	Pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", Pid)

	go signalListen()
	for {
		time.Sleep(10 * time.Second)
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileObj, _ := os.OpenFile("./tt.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer fileObj.Close()

	Rd := bufio.NewReader(fileObj)
	cont, _ := Rd.ReadSlice('#')
	fmt.Println(string(cont))

	Wr := bufio.NewWriter(fileObj)
	Wr.WriteString("WriteString writes a ## string.")
	Wr.Flush()

	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	buf := bufio.NewReaderSize(sr, 0) //默认16
	b := make([]byte, 10)

	fmt.Println("==", buf.Buffered()) // 0
	s, _ := buf.Peek(5)
	fmt.Printf("%d ==  %q\n", buf.Buffered(), s) //
	nn, er := buf.Discard(3)
	fmt.Println(nn, er)

	for n, err := 0, error(nil); err == nil; {
		fmt.Printf("Buffered:%d ==Size:%d== n:%d==  b[:n] %q ==  err:%v\n", buf.Buffered(), buf.Size(), n, b[:n], err)
		n, err = buf.Read(b)
		fmt.Printf("Buffered:%d ==Size:%d== n:%d==  b[:n] %q ==  err: %v == s: %s\n", buf.Buffered(), buf.Size(), n, b[:n], err, s)
	}

	fmt.Printf("%d ==  %q\n", buf.Buffered(), s)
}

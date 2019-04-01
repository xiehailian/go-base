package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileObj, err := os.Open("./tt.txt")
	defer fileObj.Close()

	contents, _ := ioutil.ReadAll(fileObj)  // 可读取所有io.Reader流，比如网络io
	fmt.Println(string(contents))

	if contents, _ := ioutil.ReadFile("./tt.txt"); err == nil {
		fmt.Println(string(contents))
	}

	ioutil.WriteFile("./t3.txt", contents, 0666)

}
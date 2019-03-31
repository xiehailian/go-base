package main

import "fmt"

func main() {

	switch 2 {
	case 1:
		fmt.Println("hehe")
	default:
		fmt.Println("haha")
	}

	switch a := 1; {
	case a == 1:
		fmt.Println("The integer was == 1")
		fallthrough   // 强制执行下一条case
	case a == 2:
		fmt.Println("The integer was == 2")
	case a == 3:
		fmt.Println("The integer was == 3")
		fallthrough
	case a == 4:
		fmt.Println("The integer was == 4")
	case a == 5:
		fmt.Println("The integer was == 5")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

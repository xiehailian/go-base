package main

import "fmt"

func main() {
	var a, b int = 20, 30 // 声明实际变量
	var ptra *int         // 声明指针变量
	var ptrb *int = &b

	ptra = &a // 指针变量的存储地址

	fmt.Printf("a  变量的地址是: %x\n", &a)
	fmt.Printf("b  变量的地址是: %x\n", &b)

	// 指针变量的存储地址
	fmt.Printf("ptra  变量的存储地址: %x\n", ptra)
	fmt.Printf("ptrb  变量的存储地址: %x\n", ptrb)

	// 使用指针访问值
	fmt.Printf("*ptra  变量的值: %d\n", *ptra)
	fmt.Printf("*ptrb  变量的值: %d\n", *ptrb)

	//new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体；它相当于 &T{}。
	//
	//make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel
}

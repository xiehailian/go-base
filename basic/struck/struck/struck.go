package main

import (
	"fmt"
	"reflect"
)


type Student struct {
	name string "学生名字"          // 结构体标签
	Age  int    "学生年龄"          // 结构体标签
	Room int    `json:"Roomid"` // 结构体标签
	int     // 匿名（内嵌）字段, 类型极为字段名
}

func main() {
	st := Student{"Titan", 14, 102, 11}
	fmt.Println(reflect.TypeOf(st).Field(0).Tag)
	fmt.Println(reflect.TypeOf(st).Field(1).Tag)
	fmt.Println(reflect.TypeOf(st).Field(2).Tag)
}

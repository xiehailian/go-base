package main

import (
	"fmt"
)

type Humans struct {
	name string
}

type Reader interface{
   Read()
}

// 接口内嵌接口
type Teacher interface{
  Reader
  Writer
  //Humans    // 接口不能内嵌结构体
}

type Writer interface {
	Write()
}

// 结构体内嵌接口
type Author struct {
	name string
	Writer
}

// 定义新结构体，重点是实现接口方法Write()
type Other struct {
	i int
}

func (a Author) Write() {
	fmt.Println(a.name, "  Write.")
}

// 新结构体Other实现接口方法Write()，也就可以初始化时赋值给Writer 接口
func (o Other) Write() {
	fmt.Println(" Other Write.")
}


type Human struct {
	name   string // 姓名
	Gender string // 性别
	Age    int    // 年龄
	string        // 匿名字段
}

// 结构体内嵌结构体
type Student struct {
	Human     // 匿名字段
	Room  int // 教室
	int       // 匿名字段
}

func main() {

	//  方法一：Other{99}作为Writer 接口赋值
	Ao := Author{"Other", Other{99}}
	Ao.Write()

	// 方法二：简易做法，对接口使用零值，可以完成初始化
	Au := Author{name: "Hawking"}
	Au.Write()


	//使用new方式
	stu := new(Student)
	stu.Room = 102
	stu.Human.name = "Titan"
	stu.Gender = "男"
	stu.Human.Age = 14
	stu.Human.string = "Student"

	fmt.Println("stu is:", stu)
	fmt.Printf("Student.Room is: %d\n", stu.Room)
	fmt.Printf("Student.int is: %d\n", stu.int) // 初始化时已自动给予零值：0
	fmt.Printf("Student.Human.name is: %s\n", stu.name) //  (*stu).name
	fmt.Printf("Student.Human.Gender is: %s\n", stu.Gender)
	fmt.Printf("Student.Human.Age is: %d\n", stu.Age)
	fmt.Printf("Student.Human.string is: %s\n", stu.string)

	// 使用结构体字面量赋值
	stud := Student{Room: 102, Human: Human{"Hawking", "男", 14, "Monitor"}}

	fmt.Println("stud is:", stud)
	fmt.Printf("Student.Room is: %d\n", stud.Room)
	fmt.Printf("Student.int is: %d\n", stud.int) // 初始化时已自动给予零值：0
	fmt.Printf("Student.Human.name is: %s\n", stud.Human.name)
	fmt.Printf("Student.Human.Gender is: %s\n", stud.Human.Gender)
	fmt.Printf("Student.Human.Age is: %d\n", stud.Human.Age)
	fmt.Printf("Student.Human.string is: %s\n", stud.Human.string)
}


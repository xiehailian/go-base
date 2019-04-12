package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

// a=0, b=1, c=2 可以简写为如下形式
const (
    a = iota
    b
    c
)

// a, b, c分别为0, 8, 8, 新的常量b声明后，iota 不再向下赋值
const (
    d = iota
    e = 8
    f
)

// 将内存中的字节转换为
func yoloString(b []byte) string {
	return *((*string)(unsafe.Pointer(&b)))
}


func main() {

	// string转int
	integer, _ := strconv.Atoi("134213")
	// string转int64
	integer64, _ := strconv.ParseInt("121212", 10, 64)
	//  int转string
	str := strconv.Itoa(12212)
	// int64转string
	str64 := strconv.FormatInt(23232, 10)
	fmt.Println(integer, integer64, str, str64)

	// string转为byte
	var o string = "test"
	var d []byte = []byte(o)
	var c string = string(d)
	fmt.Println(d, c)
	fmt.Println(yoloString(d))    // 用指针操作内存改变类型

	var c1 complex64 = 5 + 10i
	fmt.Printf("The value is: %v", c1)
	fmt.Println(imag(c1), real(c1))  // 获取复数虚实


	x := 1
    fmt.Println(x)     // prints 1
    // 这是可以理解为不带参数没有返回值的函数
    {
        fmt.Println(x) // prints 1
        x := 2
        fmt.Println(x) // prints 2
    }
    fmt.Println(x)     // prints 1 (不是2)


	// Go默认值传递，要引用传递就带指针
    y := [3]int{1, 2, 3}
    func(arr *[3]int) {
        (*arr)[0] = 7
        fmt.Println(arr) // prints &[7 2 3]
    }(&y)
    fmt.Println(y) // prints [7 2 3]


    // 数组的正确声明方式
    var arrAge  = [5]int{18, 20, 15, 22, 16}
	var arrName = [5]string{3: "Chris", 4: "Ron"} //指定索引位置初始化
	// {"","","","Chris","Ron"}
	var arrCount = [4]int{500, 2: 100} //指定索引位置初始化 {500,0,100,0}
	var arrLazy = [...]int{5, 6, 7, 8, 22} //数组长度初始化时根据元素多少确定
	var arrPack = [...]int{10, 5: 100} //指定索引位置初始化，数组长度与此有关 {10,0,0,0,0,100}
	var arrRoom [20]int
	var arrBed = new([20]int)    // arrBed的类型是指针*[20]int,
	fmt.Println(arrAge, arrName, arrCount, arrLazy, arrPack, arrRoom, arrBed)

	// 一个 nil 的slice中添加元素是没问题的
    var s []int
    s = append(s, 1)
	// 一个map做同样的事将会生成一个运行时的panic
    //var m map[string]int  // 错误的map定义方法
    //m["one"] = 1 //error

    // 正确的map定义方法
	map1 := make(map[string]string, 5)
	map2 := make(map[string]string)
	map3 := map[string]string{}
	map4 := map[string]string{"a": "1", "b": "2", "c": "3"}
	fmt.Println(map1, map2, map3, map4)


	// slice的使用
	sli := make([]int, 5, 10)
	fmt.Printf("切片sli长度和容量：%d, %d\n", len(sli), cap(sli))
	fmt.Println(sli)
	newsli := sli[:cap(sli)]
	fmt.Println(newsli)
	var xx = []int{2, 3, 5, 7, 11}
	fmt.Printf("切片x长度和容量：%d, %d\n", len(xx), cap(xx))
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5] // a[low : high : max]  max-low的结果表示容量  high-low为长度
	fmt.Printf("切片t长度和容量：%d, %d\n", len(t), cap(t))
	// fmt.Println(t[2]) // panic ，索引不能超过切片的长度


	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 显示: 10000 10000 数组首字节地址
	res := make([]byte, 3)
	copy(res, raw[:3]) // 利用copy 函数复制，raw 可被GC释放
}

package main

import (
	"fmt"
	"math"
)

func main() {

	// make
	a := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
	b := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
	c := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
	d := make([]int, 10, 0)         // illegal: len(s) > cap(s)
	e := make(chan int, 10)         // channel with a buffer size of 10
	f := make(map[string]int, 100)  // map with initial space for approximately 100 elements

	fmt.Println(a, b, c, d, e, f)


	// len
	const (
		c1 = imag(2i)                  // imag(2i) = 2.0 是常量
		c2 = len([10]float64{2})         // [10]float64{2} 无函数调用
		c3 = len([10]float64{c1})        // [10]float64{c1} 无函数调用
		c4 = len([10]float64{imag(2i)})   // imag(2i)常量无函数调用
		//c5 = len([10]float64{imag(z)})    // 无效: imag(z) 非常量函数调用
	)


	// append
	s0 := []int{0, 0}
	s1 := append(s0, 2)            // append 附加连接单个元素   s1 == []int{0, 0, 2}
	s2 := append(s1, 3, 5, 7)        // append 附加连接多个元素  s2 == []int{0, 0, 2, 3, 5, 7}
	s3 := append(s2, s0...)         // append 附加连接切片s0  s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
	s4 := append(s3[3:6], s3[2:]...)  // append 附加切片指定值 s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

	fmt.Println(s4)
	var s5 []interface{}
	s5 = append(s5, 42, 3.1415, "foo")  //  t == []interface{}{42, 3.1415, "foo"}

	var s6 []byte
	s6 = append(s6, "bar"...)         // append 附加连接字符串内容  b == []byte{'b', 'a', 'r' }


	// copy
	var n1 = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var n2 = make([]int, 6)
	var n3 = make([]byte, 5)
	n4 := copy(s0, n1[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
	n5 := copy(s0, s0[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
	n6 := copy(n3, "Hello, World!")  // n3 == 5, b == []byte("Hello")

	fmt.Println(n1, n2, n3, n4, n5, n6)


	// delete
	var m map[int]int
	m[1] = 2
	delete(m, 1)


	// complex real imag
	var p = complex(2, -2)             // complex128
	const n = complex(1.0, -1.4)        // 无类型complex 常量 1 - 1.4i
	x := float32(math.Cos(math.Pi/2))   // float32
	var c64 = complex(5, -x)          // complex64
	var l uint = complex(1, 0)         // 无类型 complex 常量 1 + 0i 可以转为uint
	var rl = real(c64)                // float32
	var im = imag(p)                // float64
	const o = imag(n)               // 无类型常量 -1.4

	fmt.Println(m, n, l, rl, im, o)
}

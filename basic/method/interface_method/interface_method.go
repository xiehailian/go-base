package main

import "fmt"

type T struct {
	Name string
}
type Intf interface {
	M1()
	M2()
}

func (t T) M1() {
	t.Name = "name1"
	fmt.Println("M1")
}

func (t *T) M2() {
	t.Name = "name2"
	fmt.Println("M2")
}

func main() {

	var t1 T = T{"t1"}
	t1.M1()
	t1.M2()

	//var t2 Intf = t1   // 为什么要加&
	var t2 Intf = &t1
	t2.M1()
	t2.M2()



	// interface{}(t1) 先转为空接口，再使用接口断言
	_, ok1 := interface{}(t1).(Intf)
	fmt.Println("t1 => Intf", ok1)

	_, ok2 := interface{}(t1).(T)
	fmt.Println("t1 => T", ok2)
	t1.M1()
	t1.M2()

	_, ok3 := interface{}(t1).(*T)
	fmt.Println("t1 => *T", ok3)
	t1.M1()
	t1.M2()

	_, ok4 := interface{}(&t1).(Intf)
	fmt.Println("&t1 => Intf", ok4)
	t1.M1()
	t1.M2()

	_, ok5 := interface{}(&t1).(T)
	fmt.Println("&t1 => T", ok5)

	_, ok6 := interface{}(&t1).(*T)
	fmt.Println("&t1 => *T", ok6)
	t1.M1()
	t1.M2()

}


//规则一：如果使用指针方法来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口。
//规则二：如果使用值方法来实现一个接口，那么那个类型的值和指针都能够实现对应的接口。

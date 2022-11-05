package interface_learn

import "fmt"

/*
5.2 Go语言与鸭子类型的关系
*/

type Greeting interface { // 定义接口
	Hello()
}

func Hello(one Greeting) { // 定义使用该接口为参数的函数
	one.Hello()
}

type Go struct{}

func (g Go) Hello() {
	fmt.Println("Hello, I'm Go")
}

type PHP struct{}

func (p PHP) Hello() {
	fmt.Println("Hello, I'm PHP")
}

/*
5.4 值接收者与指针接收者的区别
*/

type Person struct {
	age int
}

func (p Person) howOld() int { // 接收者是值类型的方法
	return p.age
}

func (p *Person) grouUp() { // 接收者是指针类型的方法
	p.age += 1
}

// 结论：如果实现了接收者是值类型的方法，会隐含的实现接收者是指针类型的方法；反之可能存在错误
// 区别：值类型接收者修改的都是调用对象的副本；指针类型接收者修改的都是调用对象本身
// 指针类型优点：直接修改接收者；避免复制大型结构体

/*
5.8 类型转换和断言的区别
*/

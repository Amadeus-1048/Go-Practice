package interface_learn

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	golang := Go{}
	php := PHP{}
	Hello(golang) // 隐式的将golang对象转换为Greeting类型
	Hello(php)    // 隐式的将php对象转换为Greeting类型
}

func TestPersonAge(t *testing.T) {
	// ama 是值类型
	ama := Person{age: 18}
	fmt.Println(ama.howOld()) // 值类型调用值类型的方法  传值，即使用了副本
	ama.grouUp()              // 值类型调用指针类型的方法  使用值的引用
	fmt.Println(ama.howOld())

	// deus是指针类型
	deus := &Person{age: 20}
	fmt.Println(deus.howOld()) // 指针类型调用值类型的方法  指针被解引用为值
	deus.grouUp()
	fmt.Println(deus.howOld()) // 指针类型调用指针类型的方法 传值，即复制了一份指针
}

func TestAssert(t *testing.T) {
	var i interface{} = new(Person)
	s, ok := i.(Person)
	if ok {
		fmt.Println(s)
	} else { // 需要错误处理，否则会panic
		// 这里类型断言会失败，因为 i 是 *Person 类型（new），而非Person类型
		fmt.Println("type assert error")
	}

	i = (*Person)(nil)    // 动态类型为*Person，数据为nil，其类型并不是nil
	fmt.Println(i == nil) // 与nil做比较的时候，返回为false
}

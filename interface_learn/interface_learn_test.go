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

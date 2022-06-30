package test

import (
	"fmt"
	"testing"
)

/**
鸭子类型是一种动态语言的风格，在这种风格中，一个对象有效的语义，
不是由继承自特定的类或实现特定的接口，而是由它"当前方法和属性的集合"决定。
Go 作为一种静态语言，通过接口实现了 鸭子类型，实际上是 Go 的编译器在其中作了隐匿的转换工作。
**/
type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct{}

func (g Go) sayHello() {
	fmt.Println("Hi, I am GO!")
}

type PHP struct{}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP!")
}

func TestInterface4(t *testing.T) {
	golang := Go{}
	php := PHP{}

	sayHello(golang)
	sayHello(php)
}

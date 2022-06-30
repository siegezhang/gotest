package test

import (
	"fmt"
	"testing"
	//	. "github.com/smartystreets/goconvey/convey"
)

//这里是我们普通定义的一个函数add
func add(a, b int) int {
	return a + b
}

type Int int

//这里是对Int这个自定义类型定义了一个方法add
func (i Int) add(a, b int) int {
	return a + b
}

//如果想要把计算的结果赋值给i
func (j *Int) add2(a, b int) {
	*j = Int(a + b)
}
func TestMethod(t *testing.T) {
	c := add(100, 200)
	fmt.Println(c)
	var b Int
	res := b.add(10, 100)
	fmt.Println(b)
	fmt.Println(res)

	var sum Int
	fmt.Println(sum)
	sum.add2(20, 20)
	fmt.Println(sum)
}

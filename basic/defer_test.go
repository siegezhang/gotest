package test

import (
	"fmt"
	"testing"
)

//5
func f() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

//1
func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

//5
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

//1
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
func TestDefer(t *testing.T) {
	println(f())
	println(f1())
	println(f2())
	println(f3())
}

/**
defer函数入参并不是在defer时确定不变，是值类型和引用类型的区别
*/
func TestDefer1(t *testing.T) {
	a := [3]int{1, 2, 3}
	defer fmt.Println(a)
	a[0] = 100

	b := []int{1, 2, 3}
	defer fmt.Println(b)
	b[0] = 200
}

func TestDefer2(t *testing.T) {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func TestDefer3(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	return
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}
func TestDefer4(t *testing.T) {
	fmt.Println(c())
}

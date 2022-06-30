package test

import (
	"fmt"
	"testing"
)

type A struct {
	Name string
}

type B struct {
	A
	Age int
}

//匿名字段
type C1 struct {
	int
}

func TestExtends(t *testing.T) {
	var a B
	a.Name = "siege"
	a.Age = 30
	fmt.Print(a)
	c := C1{
		12,
	}
	fmt.Print(c)
}

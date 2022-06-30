package test

import (
	"fmt"
	"testing"
)

/**
数组，'a','b','c'标识其对应的ascii的值下标
**/
func Test6(t *testing.T) {
	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	m['a'] = 3
	fmt.Println(len(m))
	fmt.Printf(`%T`, m)
}

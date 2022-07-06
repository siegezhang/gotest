package test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]string, 10)
	n := map[string]string{
		"a": "hello",
	}
	val, ok := n["a"]
	if ok {
		fmt.Print(val)
	}
	delete(n, "a")
	val1, ok1 := m["a"]
	if ok1 {
		fmt.Print(val1)
	}

}
func TestMap1(t *testing.T) {
	map1 := map[string]map[string]string{
		"key1": {
			"key1-1": "value1",
		},
		"key2": {
			"key2-1": "value2",
		},
		"key3": {
			"key3-1": "value3",
		},
	}
	for k, val := range map1 {
		fmt.Printf("key:%v,value:", k)
		for key1, val1 := range val {
			fmt.Printf("        key:%v  value:%v\n", key1, val1)
		}
	}
}

//map切片
func TestMap2(t *testing.T) {
	map1 := make([]map[string]string, 2)
	if map1[0] == nil {
		map1[0] = make(map[string]string, 2)
		map1[0]["name"] = "siege"
	}
	map2 := map[string]string{
		"name": "cage",
	}
	map1 = append(map1, map2)
	fmt.Print(map1)
}

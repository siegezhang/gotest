package test

import (
	"fmt"
	"testing"
)

/*
*
数组，'a','b','c'标识其对应的ascii的值下标
*
*/
func Test7(t *testing.T) {
	s := "你好，世界"
	fmt.Println(len(s))         // 输出 15（字节数）
	fmt.Println(len([]rune(s))) // 输出 5（字符数）

	s = "Hello, 世界"
	for _, r := range s {
		fmt.Printf("%c ", r) // 输出：H e l l o ,   世 界
	}

	s = "你好"
	runes := []rune(s)        // 将字符串转换为 rune 切片
	fmt.Printf("%U\n", runes) // 输出：[U+4F60 U+597D]

	newStr := string(runes) // 将 rune 切片转回字符串
	fmt.Println(newStr)     // 输出：你好

	s = "👋世界"
	runes = []rune(s)
	fmt.Println(len(runes)) // 输出 3（1 个 emoji + 2 个汉字）

	//按字节操作	[]byte	bytes := []byte("你好")
	//按字符操作	[]rune	runes := []rune("你好")
	//字符串底层存储	UTF-8 字节	"你好" 存储为 [0xE4 0xBD ...]
	//遍历字符串（正确）	for _, r := range s	自动解码 UTF-8，返回 rune
	//遍历字符串（按字节）	for i := 0; i < len(s); i++	返回字节值（byte）

	//rune即code point(Unicode 码点)
}

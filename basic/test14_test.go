package test

import (
	"fmt"
	"testing"
)

//接口定义
type Animal interface {
	Speak() string //定义一个方法Speak返回类型是字符串
}

//定义Dog结构
type Dog struct {
}

//定义Cat结构
type Cat struct {
}

//在结构Dog中实现接口方法
func (d Dog) Speak() string {
	return "Dog Speak"
}

//在结构Cat中实现接口方法
func (c Cat) Speak() string {
	return "Cat Speak"
}

//定义一个方法，传入参数是interface{}的集合
func PrintAll(v []interface{}) {
	//遍历对象
	for _, r := range v {
		dog, ok := r.(Dog)
		if ok {
			fmt.Println(dog.Speak())
		}
		cat, ok := r.(Cat)
		if ok {
			fmt.Println(cat.Speak())
		}
	}
}

//定义一个方法，传入参数是interface{}的集合
func PrintAnimalAll(v []Animal) {
	//遍历对象
	for _, r := range v {
		fmt.Println(r.Speak())
	}
}

func TestInterface(t *testing.T) {
	animals := []Animal{Dog{}, Cat{}}
	for _, s := range animals {
		fmt.Println(s.Speak())
	}
}

func TestInterface1(t *testing.T) {
	animals := []Animal{Dog{}, Cat{}}
	ani := make([]interface{}, len(animals), 100)
	ani[0] = Dog{}
	ani[1] = Cat{}
	PrintAll(ani)
}

func TestInterface2(t *testing.T) {
	ani := []interface{}{Dog{}, Cat{}}
	PrintAll(ani)
}

func TestInterface3(t *testing.T) {
	animals := []Animal{Dog{}, Cat{}}
	PrintAnimalAll(animals)
}

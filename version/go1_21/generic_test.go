package go1_21

import (
	"fmt"
	"testing"
)

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func TestGeneric(t *testing.T) {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}

func TestLoop(t *testing.T) {
	items := []string{"111", "211"}

	var all []*string
	for _, item := range items {
		item := item
		all = append(all, &item)
	}
	items[0] = "asd"

	for _, item := range all {
		println(*item)
	}

}

func TestLoop1(t *testing.T) {
	items := []string{"111", "211"}

	var all []*string
	for _, item := range items {
		all = append(all, &item)
	}

	for _, item := range all {
		println(*item)
	}

}
func TestLoop2(t *testing.T) {
	var prints []func()
	for _, v := range []int{1, 2, 3} {
		prints = append(prints, func() { fmt.Println(v) })
	}
	for _, print := range prints {
		print()
	}
}
func TestLoop3(t *testing.T) {
	var prints []func()
	for _, v := range []int{1, 2, 3} {
		v := v
		prints = append(prints, func() { fmt.Println(v) })
	}
	for _, print := range prints {
		print()
	}
}

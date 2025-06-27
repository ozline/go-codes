package code

/*
	go 的 set 实现，底层是个 map，看了一下八股，读写都是 O（1），性能 ok
*/

import (
	"fmt"
	"testing"
)

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

// 获取所有元素（以切片形式返回）
func (s *Set[T]) Elements() []T {
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func TestSet(t *testing.T) {

	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	fmt.Println("Set contains 2:", set.Contains(2)) // true
	fmt.Println("Set size:", set.Size())            // 3

	set.Remove(2)
	fmt.Println("Set contains 2:", set.Contains(2)) // false
	fmt.Println("Set elements:", set.Elements())    // [1 3]
}

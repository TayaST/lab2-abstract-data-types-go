package main

import "fmt"

type SNode struct {
	value string
	next  *SNode
}

type Stack struct {
	top   *SNode
	count int
}

func (s *Stack) Push(v string) {
	s.top = &SNode{value: v, next: s.top}
	s.count++
	fmt.Println("OK")
}

func (s *Stack) PushSilent(v string) {
	s.top = &SNode{value: v, next: s.top}
	s.count++
}

func (s *Stack) Pop() {
	if s.top == nil {
		fmt.Println("ERROR: stack is empty")
		return
	}
	fmt.Println(s.top.value)
	s.top = s.top.next
	s.count--
}

func (s *Stack) Peek() {
	if s.top == nil {
		fmt.Println("ERROR: stack is empty")
		return
	}
	fmt.Println(s.top.value)
}

func (s *Stack) Len() {
	fmt.Println(s.count)
}

func (s *Stack) Print() {
	fmt.Print("[ ")
	for cur := s.top; cur != nil; cur = cur.next {
		fmt.Print(cur.value, " ")
	}
	fmt.Println("]")
}

// ToSlice возвращает элементы снизу вверх (для сохранения в файл).
// Порядок: bottom → top.
func (s *Stack) ToSlice() []string {
	n := s.count
	out := make([]string, n)
	cur := s.top
	for i := n - 1; i >= 0; i-- {
		out[i] = cur.value
		cur = cur.next
	}
	return out
}
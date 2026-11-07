package main

import "fmt"

type FNode struct {
	value string
	next  *FNode
}

type ForwardList struct {
	head  *FNode
	tail  *FNode
	count int
}

// ---------- Добавление ----------

func (l *ForwardList) Push(v string) {
	n := &FNode{value: v}
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.count++
	fmt.Println("OK")
}

func (l *ForwardList) PushSilent(v string) {
	n := &FNode{value: v}
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		l.tail.next = n
		l.tail = n
	}
	l.count++
}

func (l *ForwardList) PushHead(v string) {
	n := &FNode{value: v, next: l.head}
	l.head = n
	if l.tail == nil {
		l.tail = n
	}
	l.count++
	fmt.Println("OK")
}

func (l *ForwardList) PushAfter(v, after string) {
	cur := l.head
	for cur != nil && cur.value != after {
		cur = cur.next
	}
	if cur == nil {
		fmt.Println("ERROR: not found")
		return
	}
	n := &FNode{value: v, next: cur.next}
	cur.next = n
	if cur == l.tail {
		l.tail = n
	}
	l.count++
	fmt.Println("OK")
}

func (l *ForwardList) PushBefore(v, before string) {
	if l.head == nil {
		fmt.Println("ERROR: empty list")
		return
	}
	if l.head.value == before {
		l.PushHead(v)
		return
	}
	prev := l.head
	for prev.next != nil && prev.next.value != before {
		prev = prev.next
	}
	if prev.next == nil {
		fmt.Println("ERROR: not found")
		return
	}
	n := &FNode{value: v, next: prev.next}
	prev.next = n
	l.count++
	fmt.Println("OK")
}

// ---------- Удаление ----------

func (l *ForwardList) DelFirst() {
	if l.head == nil {
		fmt.Println("ERROR: empty list")
		return
	}
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	l.count--
	fmt.Println("OK")
}

func (l *ForwardList) DelLast() {
	if l.head == nil {
		fmt.Println("ERROR: empty list")
		return
	}
	if l.head == l.tail {
		l.head, l.tail = nil, nil
		l.count--
		fmt.Println("OK")
		return
	}
	cur := l.head
	for cur.next != l.tail {
		cur = cur.next
	}
	cur.next = nil
	l.tail = cur
	l.count--
	fmt.Println("OK")
}

func (l *ForwardList) DelByValue(v string) {
	if l.head == nil {
		fmt.Println("ERROR: empty list")
		return
	}
	if l.head.value == v {
		l.DelFirst()
		return
	}
	prev := l.head
	for prev.next != nil && prev.next.value != v {
		prev = prev.next
	}
	if prev.next == nil {
		fmt.Println("ERROR: not found")
		return
	}
	if prev.next == l.tail {
		l.tail = prev
	}
	prev.next = prev.next.next
	l.count--
	fmt.Println("OK")
}

// Удалить узел ПЕРЕД узлом с индексом index.
func (l *ForwardList) DelBefore(index int) {
	if l.head == nil {
		fmt.Println("ERROR: empty list")
		return
	}
	if index <= 0 {
		fmt.Println("ERROR: index out of range")
		return
	}
	if index == 1 {
		l.DelFirst()
		return
	}
	prev := l.head
	i := 0
	for prev != nil && i < index-2 {
		prev = prev.next
		i++
	}
	if prev == nil || prev.next == nil {
		fmt.Println("ERROR: index out of range")
		return
	}
	tmp := prev.next
	prev.next = tmp.next
	if tmp == l.tail {
		l.tail = prev
	}
	l.count--
	fmt.Println("OK")
}

// Удалить узел ПОСЛЕ узла с индексом index.
func (l *ForwardList) DelAfter(index int) {
	if l.head == nil {
		fmt.Println("ERROR: empty list")
		return
	}
	cur := l.head
	i := 0
	for cur != nil && i < index {
		cur = cur.next
		i++
	}
	if cur == nil {
		fmt.Println("ERROR: index out of range")
		return
	}
	if cur.next == nil {
		fmt.Println("ERROR: nothing after this node")
		return
	}
	tmp := cur.next
	cur.next = tmp.next
	if tmp == l.tail {
		l.tail = cur
	}
	l.count--
	fmt.Println("OK")
}

// ---------- Поиск и чтение ----------

func (l *ForwardList) Get(v string) {
	cur := l.head
	idx := 0
	for cur != nil {
		if cur.value == v {
			fmt.Printf("FOUND at index %d\n", idx)
			return
		}
		cur = cur.next
		idx++
	}
	fmt.Println("NOT FOUND")
}

func (l *ForwardList) GetAt(i int) {
	cur := l.head
	j := 0
	for cur != nil && j < i {
		cur = cur.next
		j++
	}
	if cur == nil {
		fmt.Println("ERROR: index out of range")
		return
	}
	fmt.Println(cur.value)
}

func (l *ForwardList) Len() {
	fmt.Println(l.count)
}

func (l *ForwardList) Print() {
	fmt.Print("[ ")
	for cur := l.head; cur != nil; cur = cur.next {
		fmt.Print(cur.value, " ")
	}
	fmt.Println("]")
}

// ToSlice возвращает элементы от головы к хвосту (для сохранения в файл).
func (l *ForwardList) ToSlice() []string {
	out := []string{}
	for cur := l.head; cur != nil; cur = cur.next {
		out = append(out, cur.value)
	}
	return out
}
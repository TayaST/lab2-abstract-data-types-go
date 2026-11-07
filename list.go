
package main

import "fmt"

type LNode struct {
	value      string
	prev, next *LNode
}

type DoublyList struct {
	head  *LNode
	tail  *LNode
	count int
}

// ---------- Добавление ----------

func (l *DoublyList) Push(v string) {
	n := &LNode{value: v}
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
	l.count++
	fmt.Println("OK")
}

func (l *DoublyList) PushSilent(v string) {
	n := &LNode{value: v}
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
	l.count++
}

func (l *DoublyList) PushHead(v string) {
	n := &LNode{value: v}
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		n.next = l.head
		l.head.prev = n
		l.head = n
	}
	l.count++
	fmt.Println("OK")
}

func (l *DoublyList) PushAfter(v, after string) {
	cur := l.head
	for cur != nil && cur.value != after {
		cur = cur.next
	}
	if cur == nil {
		fmt.Println("ERROR: not found")
		return
	}
	n := &LNode{value: v, prev: cur, next: cur.next}
	if cur.next != nil {
		cur.next.prev = n
	} else {
		l.tail = n
	}
	cur.next = n
	l.count++
	fmt.Println("OK")
}

func (l *DoublyList) PushBefore(v, before string) {
	cur := l.head
	for cur != nil && cur.value != before {
		cur = cur.next
	}
	if cur == nil {
		fmt.Println("ERROR: not found")
		return
	}
	n := &LNode{value: v, prev: cur.prev, next: cur}
	if cur.prev != nil {
		cur.prev.next = n
	} else {
		l.head = n
	}
	cur.prev = n
	l.count++
	fmt.Println("OK")
}

// ---------- Удаление ----------

func (l *DoublyList) DelFirst() {
	if l.head == nil {
		fmt.Println("ERROR: empty")
		return
	}
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	l.count--
	fmt.Println("OK")
}

func (l *DoublyList) DelLast() {
	if l.tail == nil {
		fmt.Println("ERROR: empty")
		return
	}
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	l.count--
	fmt.Println("OK")
}

func (l *DoublyList) DelByValue(v string) {
	cur := l.head
	for cur != nil && cur.value != v {
		cur = cur.next
	}
	if cur == nil {
		fmt.Println("ERROR: not found")
		return
	}
	if cur.prev != nil {
		cur.prev.next = cur.next
	} else {
		l.head = cur.next
	}
	if cur.next != nil {
		cur.next.prev = cur.prev
	} else {
		l.tail = cur.prev
	}
	l.count--
	fmt.Println("OK")
}

// Удалить узел ПЕРЕД узлом с индексом index.
func (l *DoublyList) DelBefore(index int) {
	if l.head == nil {
		fmt.Println("ERROR: empty list")
		return
	}
	if index <= 0 {
		fmt.Println("ERROR: index out of range")
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
	tmp := cur.prev
	if tmp == nil {
		fmt.Println("ERROR: nothing before this node")
		return
	}
	if tmp.prev != nil {
		tmp.prev.next = cur
	} else {
		l.head = cur
	}
	cur.prev = tmp.prev
	l.count--
	fmt.Println("OK")
}

// Удалить узел ПОСЛЕ узла с индексом index.
func (l *DoublyList) DelAfter(index int) {
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
	tmp := cur.next
	if tmp == nil {
		fmt.Println("ERROR: nothing after this node")
		return
	}
	cur.next = tmp.next
	if tmp.next != nil {
		tmp.next.prev = cur
	} else {
		l.tail = cur
	}
	l.count--
	fmt.Println("OK")
}

// ---------- Поиск и чтение ----------

func (l *DoublyList) Get(v string) {
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

func (l *DoublyList) GetAt(i int) {
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

func (l *DoublyList) Len() {
	fmt.Println(l.count)
}

func (l *DoublyList) Print() {
	fmt.Print("[ ")
	for cur := l.head; cur != nil; cur = cur.next {
		fmt.Print(cur.value, " ")
	}
	fmt.Println("]")
}

func (l *DoublyList) PrintReverse() {
	fmt.Print("[ ")
	for cur := l.tail; cur != nil; cur = cur.prev {
		fmt.Print(cur.value, " ")
	}
	fmt.Println("]")
}

// ToSlice возвращает элементы от головы к хвосту (для сохранения в файл).
func (l *DoublyList) ToSlice() []string {
	out := []string{}
	for cur := l.head; cur != nil; cur = cur.next {
		out = append(out, cur.value)
	}
	return out
}
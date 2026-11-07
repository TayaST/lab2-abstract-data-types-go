package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const TableSize = 10007

type HashSet struct {
	table [TableSize]ForwardList
}

func (s *HashSet) Init() {
	for i := 0; i < TableSize; i++ { s.table[i].head = nil; s.table[i].tail = nil }
}

func (s *HashSet) hash(key string) int {
	h := 5381
	for _, c := range key { h = ((h << 5) + h) + int(c) }
	return h % TableSize
}

func (s *HashSet) Add(key string) {
	idx := s.hash(key)
	cur := s.table[idx].head
	for cur != nil {
		if cur.value == key { return }
		cur = cur.next
	}
	s.table[idx].PushSilent(key)
}

func (s *HashSet) Del(key string) {
	idx := s.hash(key)
	cur := s.table[idx].head
	var prev *FNode
	for cur != nil {
		if cur.value == key {
			if prev == nil { s.table[idx].head = cur.next } else { prev.next = cur.next }
			if cur == s.table[idx].tail { s.table[idx].tail = prev }
			return
		}
		prev = cur
		cur = cur.next
	}
}

func (s *HashSet) Contains(key string) bool {
	idx := s.hash(key)
	cur := s.table[idx].head
	for cur != nil {
		if cur.value == key { return true }
		cur = cur.next
	}
	return false
}

func (s *HashSet) Print() {
	fmt.Print("{")
	first := true
	for i := 0; i < TableSize; i++ {
		for cur := s.table[i].head; cur != nil; cur = cur.next {
			if !first { fmt.Print(", ") }
			fmt.Print(cur.value)
			first = false
		}
	}
	fmt.Println("}")
}

func runSets() {
	fmt.Println("\n=== Задание 2: Множество ===")
	var s HashSet
	s.Init()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() { break }
		line := scanner.Text()
		if line == "exit" { break }
		parts := strings.Fields(line)
		if len(parts) < 2 { continue }

		cmd, arg := parts[0], parts[1]
		switch cmd {
		case "SETADD": s.Add(arg); fmt.Println("OK")
		case "SETDEL": s.Del(arg); fmt.Println("OK")
		case "SET_AT":
			if s.Contains(arg) { fmt.Println("YES") } else { fmt.Println("NO") }
		}
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const HASH_SIZE = 20011

type DoubleHash struct {
	keys [HASH_SIZE]string
	used [HASH_SIZE]bool
}

func (h *DoubleHash) hash1(s string) int {
	res := 5381
	for _, c := range s {
		res = ((res << 5) + res) + int(c)
		res = res & 0x7FFFFFFF
	}
	return res % HASH_SIZE
}

func (h *DoubleHash) hash2(s string) int {
	res := 0
	for _, c := range s {
		res = res*31 + int(c)
		res = res & 0x7FFFFFFF
	}
	step := 1 + (res % (HASH_SIZE - 1))
	if step == 0 {
		step = 1
	}
	return step
}

func (h *DoubleHash) Insert(key string) {
	idx := h.hash1(key)
	step := h.hash2(key)
	
	for attempts := 0; attempts < HASH_SIZE; attempts++ {
		if !h.used[idx] {
			h.keys[idx] = key
			h.used[idx] = true
			return
		}
		if h.keys[idx] == key {
			return
		}
		idx = (idx + step) % HASH_SIZE
	}
}

func runHash() {
	fmt.Println("\n=== Задание 6: Родословная (Вариант 9) ===")
	
	ht := &DoubleHash{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("N: ")
	var n int
	fmt.Scan(&n)
	reader.ReadString('\n') // Очистка буфера от Enter

	type Person struct {
		name   string
		parent int
		desc   int
	}
	
	people := make([]Person, 0)
	nameToID := make(map[string]int)

	getID := func(name string) int {
		if id, ok := nameToID[name]; ok {
			return id
		}
		ht.Insert(name) // Используем нашу хеш-таблицу
		id := len(people)
		people = append(people, Person{name: name, parent: -1})
		nameToID[name] = id
		return id
	}

	// Читаем пары
	for i := 0; i < n-1; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		c := getID(parts[0])
		p := getID(parts[1])
		people[c].parent = p
	}

	// Итеративный подсчет ВСЕХ потомков (через очередь)
	for i := range people {
		count := 0
		queue := []int{i}
		
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			
			for j := range people {
				if people[j].parent == curr {
					count++
					queue = append(queue, j) // Добавляем детей в очередь, чтобы найти их детей
				}
			}
		}
		people[i].desc = count
	}

	// Сортировка по алфавиту
	sort.Slice(people, func(i, j int) bool {
		return people[i].name < people[j].name
	})

	fmt.Print("\nВывод:\n")
	for _, p := range people {
		fmt.Printf("%s %d\n", p.name, p.desc)
	}
}
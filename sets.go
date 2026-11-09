package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// HashSet - структура для нашего множества
type HashSet struct {
	data map[string]bool
}

// Инициализация множества
func (s *HashSet) Init() {
	s.data = make(map[string]bool)
}

// Добавление элемента (с проверкой на дубликат)
func (s *HashSet) Add(val string) bool {
	if s.data[val] {
		return false // Уже есть
	}
	s.data[val] = true
	return true
}

// Удаление элемента
func (s *HashSet) Del(val string) {
	delete(s.data, val)
}

// Проверка наличия элемента
func (s *HashSet) Contains(val string) bool {
	return s.data[val]
}

// Вывод множества
func (s *HashSet) Print() {
	fmt.Print("{")
	first := true
	for k := range s.data {
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(k)
		first = false
	}
	fmt.Println("}")
}

// Обработка одной команды
func processSetQuery(s *HashSet, query string) bool { // Возвращает false, если надо выйти
	parts := strings.Fields(query)
	if len(parts) == 0 {
		return true
	}

	cmd := parts[0]
	
	if cmd == "SETADD" && len(parts) >= 2 {
		if s.Add(parts[1]) {
			fmt.Println("OK")
		} else {
			fmt.Println("ERROR: element already exists")
		}
	} else if cmd == "SETDEL" && len(parts) >= 2 {
		s.Del(parts[1])
		fmt.Println("OK")
	} else if cmd == "SET_AT" && len(parts) >= 2 {
		if s.Contains(parts[1]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else if cmd == "exit" || cmd == "EXIT" {
		return false
	} else {
		fmt.Println("ERROR: unknown command")
	}
	
	return true
}

func runSets() {
	fmt.Println("\n=== Задание 2: АТД Множество (Интерактив) ===")
	fmt.Println("Команды: SETADD x, SETDEL x, SET_AT x")
	fmt.Println("Для выхода: exit")

	var s HashSet
	s.Init()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		
		if !processSetQuery(&s, line) {
			break
		}
	}
}
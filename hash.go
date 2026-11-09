package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

const TABLE_SIZE = 20011

// ==========================================
// 6.1. ХЕШ-ТАБЛИЦА С ЛИНЕЙНЫМ ПРОБИРОВАНИЕМ
// ==========================================
type LinearHash struct {
	keys        [TABLE_SIZE]string
	values      [TABLE_SIZE]int
	isOccupied  [TABLE_SIZE]bool
	isDeleted   [TABLE_SIZE]bool
	collisions  int
}

func (h *LinearHash) Init() {
	for i := 0; i < TABLE_SIZE; i++ {
		h.isOccupied[i] = false
		h.isDeleted[i] = false
		h.values[i] = 0
	}
	h.collisions = 0
}

func (h *LinearHash) Hash(str string) int {
	var hash uint64 = 5381
	for _, c := range str {
		hash = ((hash << 5) + hash) + uint64(c)
	}
	return int(hash % TABLE_SIZE)
}

func (h *LinearHash) Insert(key string, value int) bool {
	idx := h.Hash(key)
	firstDeleted := -1
	collided := false

	for i := 0; i < TABLE_SIZE; i++ {
		pos := (idx + i) % TABLE_SIZE

		if h.isOccupied[pos] && h.keys[pos] == key {
			h.values[pos] = value
			return true
		}
		if !h.isOccupied[pos] && h.isDeleted[pos] && firstDeleted == -1 {
			firstDeleted = pos
		}
		if !h.isOccupied[pos] && !h.isDeleted[pos] {
			target := pos
			if firstDeleted != -1 {
				target = firstDeleted
			}
			h.keys[target] = key
			h.values[target] = value
			h.isOccupied[target] = true
			h.isDeleted[target] = false
			if collided {
				h.collisions++
			}
			return true
		}
		collided = true
	}

	if firstDeleted != -1 {
		h.keys[firstDeleted] = key
		h.values[firstDeleted] = value
		h.isOccupied[firstDeleted] = true
		h.isDeleted[firstDeleted] = false
		if collided {
			h.collisions++
		}
		return true
	}
	return false
}

func (h *LinearHash) Search(key string) bool {
	idx := h.Hash(key)
	for i := 0; i < TABLE_SIZE; i++ {
		pos := (idx + i) % TABLE_SIZE
		if !h.isOccupied[pos] && !h.isDeleted[pos] {
			return false
		}
		if h.isOccupied[pos] && h.keys[pos] == key {
			return true
		}
	}
	return false
}

func (h *LinearHash) Remove(key string) bool {
	idx := h.Hash(key)
	for i := 0; i < TABLE_SIZE; i++ {
		pos := (idx + i) % TABLE_SIZE
		if !h.isOccupied[pos] && !h.isDeleted[pos] {
			return false
		}
		if h.isOccupied[pos] && h.keys[pos] == key {
			h.isOccupied[pos] = false
			h.isDeleted[pos] = true
			return true
		}
	}
	return false
}

// ==========================================
// 6.1. ХЕШ-ТАБЛИЦА С ДВОЙНЫМ ХЕШИРОВАНИЕМ
// ==========================================
type DoubleHash struct {
	keys       [TABLE_SIZE]string
	values     [TABLE_SIZE]int
	isOccupied [TABLE_SIZE]bool
	isDeleted  [TABLE_SIZE]bool
	collisions int
}

func (h *DoubleHash) Init() {
	for i := 0; i < TABLE_SIZE; i++ {
		h.isOccupied[i] = false
		h.isDeleted[i] = false
		h.values[i] = 0
	}
	h.collisions = 0
}

func (h *DoubleHash) Hash1(str string) int {
	var hash uint64 = 5381
	for _, c := range str {
		hash = ((hash << 5) + hash) + uint64(c)
	}
	return int(hash % TABLE_SIZE)
}

func (h *DoubleHash) Hash2(str string) int {
	var hash uint64 = 0
	for _, c := range str {
		hash = (hash * 31) + uint64(c)
	}
	return 1 + int(hash%(TABLE_SIZE-1))
}

func (h *DoubleHash) Insert(key string, value int) bool {
	idx := h.Hash1(key)
	step := h.Hash2(key)
	firstDeleted := -1
	collided := false

	for i := 0; i < TABLE_SIZE; i++ {
		pos := (idx + i*step) % TABLE_SIZE

		if h.isOccupied[pos] && h.keys[pos] == key {
			h.values[pos] = value
			return true
		}
		if !h.isOccupied[pos] && h.isDeleted[pos] && firstDeleted == -1 {
			firstDeleted = pos
		}
		if !h.isOccupied[pos] && !h.isDeleted[pos] {
			target := pos
			if firstDeleted != -1 {
				target = firstDeleted
			}
			h.keys[target] = key
			h.values[target] = value
			h.isOccupied[target] = true
			h.isDeleted[target] = false
			if collided {
				h.collisions++
			}
			return true
		}
		collided = true
	}

	if firstDeleted != -1 {
		h.keys[firstDeleted] = key
		h.values[firstDeleted] = value
		h.isOccupied[firstDeleted] = true
		h.isDeleted[firstDeleted] = false
		if collided {
			h.collisions++
		}
		return true
	}
	return false
}

func (h *DoubleHash) Search(key string) bool {
	idx := h.Hash1(key)
	step := h.Hash2(key)
	for i := 0; i < TABLE_SIZE; i++ {
		pos := (idx + i*step) % TABLE_SIZE
		if !h.isOccupied[pos] && !h.isDeleted[pos] {
			return false
		}
		if h.isOccupied[pos] && h.keys[pos] == key {
			return true
		}
	}
	return false
}

func (h *DoubleHash) Remove(key string) bool {
	idx := h.Hash1(key)
	step := h.Hash2(key)
	for i := 0; i < TABLE_SIZE; i++ {
		pos := (idx + i*step) % TABLE_SIZE
		if !h.isOccupied[pos] && !h.isDeleted[pos] {
			return false
		}
		if h.isOccupied[pos] && h.keys[pos] == key {
			h.isOccupied[pos] = false
			h.isDeleted[pos] = true
			return true
		}
	}
	return false
}

// ==========================================
// ТЕСТ ДЛЯ ПОДЗАДАНИЯ 6.1
// ==========================================
func testHashMethods() {
	fmt.Println("\n--- Эмпирический анализ 6.1 ---")
	
	lh := &LinearHash{}
	dh := &DoubleHash{}
	lh.Init()
	dh.Init()

	N := 5000
	testStrings := make([]string, N)
	for i := 0; i < N; i++ {
		testStrings[i] = fmt.Sprintf("key_%d_%d", i, i%1000)
	}

	start := time.Now()
	for i := 0; i < N; i++ {
		lh.Insert(testStrings[i], i)
	}
	lhTime := time.Since(start).Milliseconds()

	start = time.Now()
	for i := 0; i < N; i++ {
		dh.Insert(testStrings[i], i)
	}
	dhTime := time.Since(start).Milliseconds()

	found1 := lh.Search(testStrings[0])
	found2 := dh.Search(testStrings[N-1])

	fmt.Printf("Вставка %d элементов:\n", N)
	fmt.Printf("1) Линейное пробирование:\n   Время: %d мс\n   Коллизий: %d\n", lhTime, lh.collisions)
	fmt.Printf("2) Двойное хеширование:\n   Время: %d мс\n   Коллизий: %d\n", dhTime, dh.collisions)

	fmt.Println("\nПроверка поиска:")
	fmt.Printf("   Линейное:   key[0] -> %s\n", map[bool]string{true: "найден", false: "не найден"}[found1])
	fmt.Printf("   Двойное:    key[N-1] -> %s\n", map[bool]string{true: "найден", false: "не найден"}[found2])

	removed := lh.Remove(testStrings[0])
	stillThere := lh.Search(testStrings[0])
	
	fmt.Println("\nПроверка удаления (линейное):")
	fmt.Printf("   remove(key[0]): %s\n", map[bool]string{true: "OK", false: "не найден"}[removed])
	fmt.Printf("   search(key[0]) после удаления: %s\n", map[bool]string{true: "всё ещё найден (ошибка!)", false: "не найден (корректно)"}[stillThere])
	
	fmt.Println("\nВывод: Двойное хеширование обычно дает меньше коллизий.")
	fmt.Println("------------------------------")
}

// ==========================================
// 6.2. ЗАДАЧА "РОДОСЛОВНАЯ"
// ==========================================
type Person struct {
	name             string
	parentIdx        int
	descendantsCount int
}

func countDescendants(pIdx int, people []Person) int {
	count := 0
	for i := 0; i < len(people); i++ {
		curr := people[i].parentIdx
		steps := 0
		for curr != -1 && steps <= len(people) {
			if curr == pIdx {
				count++
				break
			}
			curr = people[curr].parentIdx
			steps++
		}
	}
	return count
}

func runHash() {
	testHashMethods()

	fmt.Println("\n=== Задание 6.2: Родословная (Вариант 9) ===")
	
	ht := &DoubleHash{}
	ht.Init()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите число элементов (N): ")
	var n int
	fmt.Scan(&n)
	reader.ReadString('\n') // Очистка буфера

	people := make([]Person, 0)

	getPersonIdx := func(name string) int {
		// Ищем в массиве
		for i := 0; i < len(people); i++ {
			if people[i].name == name {
				return i
			}
		}
		// Добавляем в хеш-таблицу
		ht.Insert(name, len(people))
		people = append(people, Person{name: name, parentIdx: -1, descendantsCount: 0})
		return len(people) - 1
	}

	for i := 0; i < n-1; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		cIdx := getPersonIdx(parts[0])
		pIdx := getPersonIdx(parts[1])
		people[cIdx].parentIdx = pIdx
	}

	// Считаем потомков
	for i := 0; i < len(people); i++ {
		people[i].descendantsCount = countDescendants(i, people)
	}

	// Сортируем по алфавиту
	sort.Slice(people, func(i, j int) bool {
		return people[i].name < people[j].name
	})

	fmt.Println("\nВывод:")
	for _, p := range people {
		fmt.Printf("%s %d\n", p.name, p.descendantsCount)
	}
}
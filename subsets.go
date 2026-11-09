package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1. Объединение (A U B)
func setUnion(a, b *HashSet) *HashSet {
	res := &HashSet{}
	res.Init()
	// Добавляем все элементы из A
	for k := range a.data {
		res.Add(k)
	}
	// Добавляем все элементы из B
	for k := range b.data {
		res.Add(k)
	}
	return res
}

// 2. Пересечение (A ∩ B)
func setIntersection(a, b *HashSet) *HashSet {
	res := &HashSet{}
	res.Init()
	for k := range a.data {
		if b.Contains(k) {
			res.Add(k)
		}
	}
	return res
}

// 3. Разность (A \ B)
func setDifference(a, b *HashSet) *HashSet {
	res := &HashSet{}
	res.Init()
	for k := range a.data {
		if !b.Contains(k) {
			res.Add(k)
		}
	}
	return res
}

func runSubsets() {
	fmt.Println("\n=== Задание 3: Операции с множествами ===")
	a, b := &HashSet{}, &HashSet{}
	a.Init()
	b.Init()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите элементы A (через пробел):")
	lineA, _ := reader.ReadString('\n')
	lineA = strings.TrimSpace(lineA)
	for _, s := range strings.Fields(lineA) {
		a.Add(s)
	}

	fmt.Println("Введите элементы B (через пробел):")
	lineB, _ := reader.ReadString('\n')
	lineB = strings.TrimSpace(lineB)
	for _, s := range strings.Fields(lineB) {
		b.Add(s)
	}

	fmt.Print("A = "); a.Print()
	fmt.Print("B = "); b.Print()

	fmt.Print("Объединение: "); setUnion(a, b).Print()
	fmt.Print("Пересечение: "); setIntersection(a, b).Print()
	fmt.Print("Разность: "); setDifference(a, b).Print()
}
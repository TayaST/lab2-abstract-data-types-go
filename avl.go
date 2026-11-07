package main

import "fmt"

// Вычисление высоты поддерева
func getHeight(n *TNode) int {
	if n == nil {
		return 0
	}
	lh, rh := getHeight(n.left), getHeight(n.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

// Проверка на АВЛ-сбалансированность
func isAVL(n *TNode) bool {
	if n == nil {
		return true
	}
	lh, rh := getHeight(n.left), getHeight(n.right)
	diff := lh - rh
	if diff < 0 {
		diff = -diff
	}

	// Если разница больше 1, дерево не сбалансировано
	if diff > 1 {
		return false
	}
	
	// Рекурсивно проверяем левое и правое поддерево
	return isAVL(n.left) && isAVL(n.right)
}

func runTree() {
	fmt.Println("\n=== Задание 5: АВЛ-сбалансированность (Вариант 3) ===")
	t := &BST{}
	fmt.Println("Введите числа (0 для завершения):")
	for {
		var v int
		_, err := fmt.Scan(&v)
		if err != nil || v == 0 {
			break
		}
		t.InsertSilent(v) // Используем тихую вставку из твоего tree.go
	}

	if isAVL(t.root) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
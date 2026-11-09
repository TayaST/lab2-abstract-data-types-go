package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isMatch(s, p string) bool {
	n, m := len(s), len(p)
	// Создаем двумерный срез (таблицу) для DP
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	
	// Базовый случай: пустая строка соответствует пустому шаблону
	dp[0][0] = true

	// Инициализация для шаблона, начинающегося со звездочек
	for j := 1; j <= m; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	// Заполняем таблицу
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '*' {
				// Звездочка может соответствовать пустой строке или поглотить символ
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				// Совпадение или '?'
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[n][m]
}

func runWildcard() {
	fmt.Println("\n=== Задание 4: Сопоставление с паттерном (Вариант 3) ===")
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите последовательность символов: ")
	sequence, _ := reader.ReadString('\n')
	sequence = strings.TrimSpace(sequence)

	fmt.Print("Введите шаблон (можно использовать '?' и '*'): ")
	pattern, _ := reader.ReadString('\n')
	pattern = strings.TrimSpace(pattern)

	if isMatch(sequence, pattern) {
		fmt.Println("\nРезультат: Соответствует")
	} else {
		fmt.Println("\nРезультат: Не соответствует")
	}
}
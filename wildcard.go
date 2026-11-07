package main

import "fmt"

func isMatch(s, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := range dp { dp[i] = make([]bool, m+1) }
	dp[0][0] = true

	for j := 1; j <= m; j++ {
		if p[j-1] == '*' { dp[0][j] = dp[0][j-1] }
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}

func runWildcard() {
	fmt.Println("\n=== Задание 4: Wildcard ===")
	var s, p string
	fmt.Print("Строка: "); fmt.Scanln(&s)
	fmt.Print("Шаблон: "); fmt.Scanln(&p)

	if isMatch(s, p) {
		fmt.Println("Соответствует")
	} else {
		fmt.Println("Не соответствует")
	}
}
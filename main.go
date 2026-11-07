package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // <--- Единый ридер

	for {
		fmt.Println("\n=== Лабораторная работа 2 ===")
		fmt.Println("1. XML")
		fmt.Println("2. Множество")
		fmt.Println("3. Операции с множеством")
		fmt.Println("4. Wildcard")
		fmt.Println("5. Дерево (AVL)")
		fmt.Println("6. Родословная")
		fmt.Println("7. LRU Кэш")
		fmt.Println("0. Выход")
		fmt.Print("Выбор: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "" { continue } // Игнорируем пустой Enter

		switch choice {
		case "0": 
			return
		case "1": 
			runXML() // Убедись, что runXML тоже использует reader, либо пересоздает его
		case "2": 
			runSets()
		case "3": 
			runSubsets()
		case "4": 
			runWildcard()
		case "5": 
			runTree()
		case "6": 
			runHash()
		case "7": 
			runCache()
		default: 
			fmt.Println("Неверный выбор")
		}
	}
}
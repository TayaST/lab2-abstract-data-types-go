package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция проверки корректности XML-строки
func isValidXML(s string) bool {
	st := &Stack{}
	
	// Идем по строке посимвольно
	for i := 0; i < len(s); i++ {
		if s[i] == '<' {
			// Ищем закрывающую скобку '>'
			end := strings.IndexByte(s[i:], '>')
			if end == -1 { 
				return false 
			}
			
			// Извлекаем содержимое тега
			tag := s[i+1 : i+end]
			tag = strings.ReplaceAll(tag, " ", "") // Убираем пробелы, если есть

			if len(tag) == 0 { 
				return false 
			}

			if tag[0] == '/' {
				// Закрывающий тег
				name := tag[1:]
				// Если стек пуст или имя в стеке не совпадает - ошибка
				if st.top == nil || st.top.value != name {
					return false
				}
				// Извлекаем из стека (тихий pop)
				st.top = st.top.next 
				st.count--
			} else {
				// Открывающий тег (тихий push в стек)
				st.top = &SNode{value: tag, next: st.top}
				st.count++
			}
			// Перепрыгиваем через тег
			i += end
		} else if s[i] == '>' {
			// Если встретили '>' вне тега - ошибка
			return false
		}
	}
	
	// В конце стек должен быть пуст
	return st.top == nil
}

func runXML() {
	fmt.Println("\n=== Задание 1: XML (Вариант 7) ===")
	fmt.Print("Введите поврежденную строку: ")
	
	// Читаем всю строку целиком, включая пробелы
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Убираем \n и \r в конце

	if input == "" {
		return
	}

	// Алфавит для замены: строчные буквы, <, >, /
	alphabet := "abcdefghijklmnopqrstuvwxyz<>/"
	
	// Перебираем каждый символ в строке
	for i := 0; i < len(input); i++ {
		orig := input[i]
		
		// Пробуем заменить на каждый символ из алфавита
		for _, c := range []byte(alphabet) {
			if c == orig { 
				continue 
			}
			
			// Создаем копию строки с замененным символом
			temp := []byte(input)
			temp[i] = c
			newStr := string(temp)
			
			// Если строка стала корректной - выводим и завершаем
			if isValidXML(newStr) {
				fmt.Printf("Исправленная строка: %s\n", newStr)
				return
			}
		}
	}
	
	// Если ни одна замена не подошла
	fmt.Println("Incorrect")
}

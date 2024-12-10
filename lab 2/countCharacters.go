package main

import "fmt"

// countCharacters подсчитывает количество вхождений каждого символа в строке
func countCharacters(s string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	return charCount
}

func main() {
	// Пример использования функции countCharacters
	text := "hello, world!"
	result := countCharacters(text)
	fmt.Println(result) // Результат: карта с количеством вхождений символов
}

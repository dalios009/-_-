package main

import (
	"errors"
	"fmt"
)

// listEven возвращает срез чётных чисел в заданном диапазоне и ошибку, если границы неверны
func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница больше правой")
	}

	var evens []int
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens, nil
}

func main() {
	// Пример использования функции listEven
	evens, err := listEven(10, 20)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Чётные числа:", evens) // Результат: Чётные числа: [10 12 14 16 18 20]
	}
}

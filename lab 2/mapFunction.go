package main

import "fmt"

// Map применяет переданную функцию к каждому элементу среза
func Map(slice []float64, fn func(float64) float64) []float64 {
	result := make([]float64, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func main() {
	// Пример использования функции Map
	slice := []float64{1, 2, 3, 4, 5}

	// Анонимная функция для возведения в квадрат
	square := func(x float64) float64 {
		return x * x
	}

	// Применение функции к каждому элементу среза
	newSlice := Map(slice, square)

	fmt.Println("Исходный срез:", slice)           // Результат: Исходный срез: [1 2 3 4 5]
	fmt.Println("После применения Map:", newSlice) // Результат: После применения Map: [1 4 9 16 25]
}

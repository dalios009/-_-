package main

import (
	"fmt"
	"math"
)

// Структура Point описывает точку с координатами X и Y
type Point struct {
	X, Y float64
}

// Структура Line описывает отрезок с начальной и конечной точкой
type Line struct {
	Start, End Point
}

// Метод Length вычисляет длину отрезка
func (l Line) Length() float64 {
	dx := l.End.X - l.Start.X
	dy := l.End.Y - l.Start.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Пример использования структуры Line и метода Length
	p1 := Point{0, 0}
	p2 := Point{3, 4}
	line := Line{p1, p2}

	fmt.Printf("Длина отрезка: %.2f\n", line.Length()) // Результат: Длина отрезка: 5.00
}

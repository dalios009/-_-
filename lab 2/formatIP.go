package main

import "fmt"

// formatIP принимает массив из 4 байтов и возвращает строку IP-адреса
func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	// Пример использования функции formatIP
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println(formatIP(ip)) // Результат: 127.0.0.1
}

package main

import (
	"fmt"
	"sync"
)

// Function to process numbers from the channel
func count(numbers chan int, wg *sync.WaitGroup) {
	// Ensure Done() is called when this goroutine finishes
	defer wg.Done()

	// Loop to read from the channel
	for num := range numbers {
		// Perform the operation (squaring the number)
		squared := num * num
		// Print the result
		fmt.Println("Processed:", squared)
	}
}

func main() {
	// Create a channel for sending numbers
	numbers := make(chan int)

	// Use WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Start the `count` function in a new goroutine
	wg.Add(1) // Indicate that we're waiting for 1 goroutine
	go count(numbers, &wg)

	// Send some numbers into the channel
	for i := 1; i <= 5; i++ {
		fmt.Println("Sending:", i)
		numbers <- i
	}

	// Close the channel to indicate no more values will be sent
	close(numbers)

	// Wait for all goroutines to finish
	wg.Wait()

	// Optional: Sleep for a short time to ensure all goroutines finish printing
	// time.Sleep(1 * time.Second)
}

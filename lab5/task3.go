package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
)

func loadPNG1(filePath string) (image.Image, error) {
	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the image as PNG
	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode PNG: %v", err)
	}

	// Return the decoded image
	return img, nil
}

func applyGrayscaleParallel(img draw.Image, startY, endY int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup when this goroutine finishes

	bounds := img.Bounds()

	// Loop through each pixel in the assigned rows
	for y := startY; y < endY; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the original color of the pixel
			r, g, b, _ := img.At(x, y).RGBA()

			// Calculate the grayscale value (average of RGB)
			gray := uint8((r + g + b) / 3 >> 8)

			// Set the new color (grayscale)
			newColor := color.RGBA{R: gray, G: gray, B: gray, A: 255}
			img.Set(x, y, newColor)
		}
	}
}

func main() {
	// Load the PNG image
	filePath := "C:/Users/dalios009/Desktop/bvt2204/input.png"
	img, err := loadPNG1(filePath)
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	// Ensure the image is editable
	drawImg, ok := img.(draw.Image)
	if !ok {
		fmt.Println("Image is not editable")
		return
	}

	// Get image bounds
	bounds := drawImg.Bounds()
	height := bounds.Dy()

	// Divide the work into parts
	numGoroutines := 4 // Adjust the number of goroutines as needed
	rowsPerGoroutine := height / numGoroutines

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		startY := i * rowsPerGoroutine
		endY := startY + rowsPerGoroutine

		// Ensure the last goroutine processes all remaining rows
		if i == numGoroutines-1 {
			endY = height
		}

		wg.Add(1)
		go applyGrayscaleParallel(drawImg, startY, endY, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Save the modified image
	outputFile, err := os.Create("C:/Users/dalios009/Desktop/bvt2204/output_parallel.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Encode the modified image as PNG
	err = png.Encode(outputFile, drawImg)
	if err != nil {
		fmt.Println("Error saving modified image:", err)
		return
	}

	fmt.Println("Parallel image processing completed. Output saved as 'output_parallel.png'")
}

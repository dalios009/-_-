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

func loadPNG2(filePath string) (image.Image, error) {
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

func applyConvolution(img draw.Image, kernel [][]int, kernelSum int, startY, endY int, result *image.RGBA, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup when this goroutine finishes

	bounds := img.Bounds()

	// Iterate through each pixel in the assigned rows
	for y := startY; y < endY; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Apply the convolution kernel
			var rSum, gSum, bSum int
			kernelSize := len(kernel)

			for ky := 0; ky < kernelSize; ky++ {
				for kx := 0; kx < kernelSize; kx++ {
					// Get the coordinates of the neighboring pixel
					nx := x + kx - kernelSize/2
					ny := y + ky - kernelSize/2

					// Check if the neighbor is within bounds
					if nx >= bounds.Min.X && nx < bounds.Max.X && ny >= bounds.Min.Y && ny < bounds.Max.Y {
						r, g, b, _ := img.At(nx, ny).RGBA()

						// Accumulate the weighted values
						weight := kernel[ky][kx]
						rSum += int(r>>8) * weight
						gSum += int(g>>8) * weight
						bSum += int(b>>8) * weight
					}
				}
			}

			// Normalize the accumulated values
			if kernelSum > 0 {
				rSum /= kernelSum
				gSum /= kernelSum
				bSum /= kernelSum
			}

			// Clamp the values to the range [0, 255]
			rSum = clamp(rSum, 0, 255)
			gSum = clamp(gSum, 0, 255)
			bSum = clamp(bSum, 0, 255)

			// Set the new pixel value in the result image
			newColor := color.RGBA{R: uint8(rSum), G: uint8(gSum), B: uint8(bSum), A: 255}
			result.Set(x, y, newColor)
		}
	}
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func main() {
	// Load the PNG image
	filePath := "C:/Users/dalios009/Desktop/bvt2204/input.png"
	img, err := loadPNG2(filePath)
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

	// Define the Gaussian blur kernel
	kernel := [][]int{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}
	kernelSum := 16 // Sum of all weights in the kernel

	// Create a new RGBA image to store the result
	bounds := drawImg.Bounds()
	result := image.NewRGBA(bounds)

	// Divide the work into parts
	numGoroutines := 4 // Adjust based on your system
	rowsPerGoroutine := bounds.Dy() / numGoroutines

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		startY := i * rowsPerGoroutine
		endY := startY + rowsPerGoroutine

		// Ensure the last goroutine processes all remaining rows
		if i == numGoroutines-1 {
			endY = bounds.Max.Y
		}

		wg.Add(1)
		go applyConvolution(drawImg, kernel, kernelSum, startY, endY, result, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Save the modified image
	outputFile, err := os.Create("C:/Users/dalios009/Desktop/bvt2204/output_convolution.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Encode the modified image as PNG
	err = png.Encode(outputFile, result)
	if err != nil {
		fmt.Println("Error saving modified image:", err)
		return
	}

	fmt.Println("Convolution filter applied successfully. Output saved as 'output_convolution.png'")
}

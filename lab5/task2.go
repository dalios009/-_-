package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func loadPNG(filePath string) (image.Image, error) {
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

func applyGrayscale(img draw.Image) {
	// Get image bounds
	bounds := img.Bounds()

	// Loop through each pixel
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
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
	img, err := loadPNG(filePath)
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

	// Apply grayscale filter
	applyGrayscale(drawImg)

	// Save the modified image
	outputFile, err := os.Create("C:/Users/dalios009/Desktop/bvt2204/output.png")
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

	fmt.Println("Image processing completed. Output saved as 'output.png'")
}

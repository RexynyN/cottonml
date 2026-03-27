package main

import (
	"fmt"
	"image"

	// Import specific image format packages for their side effects
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

var (
	whitelist = []string{".jpg", ".png", ".jpeg", ".gif", ".jfif"}
)

func Color2RGBA(color color.Color) (r uint8, g uint8, b uint8, a uint8) {
	r1, g1, b1, a1 := color.RGBA()

	r = uint8(r1 >> 8)
	g = uint8(g1 >> 8)
	b = uint8(b1 >> 8)
	a = uint8(a1 >> 8)
	return
}

// func main() {

// 	img, format, err := readImageFile(filePath)
// 	if err != nil {
// 		log.Fatalf("failed to read image file: %v", err)
// 	}

// 	fmt.Printf("Successfully read image (format: %s, dimensions: %dx%d)\n", format, img.Bounds().Dx(), img.Bounds().Dy())

// 	r, g, b, a := Color2RGBA(img.At(34, 35))
// 	fmt.Printf("Pixel (RGBA): %d, %d, %d, %d\n", r, g, b, a)
// 	fmt.Println(r, g, b, a)
// }

// readImageFile opens a file and decodes it into an image.Image
func readImageFile(filePath string) (image.Image, string, error) {
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("error opening file: %w", err)
	}
	// Ensure the file is closed after the function returns
	defer f.Close()

	// Decode the image. The image.Decode function automatically detects the format
	// based on the imported format packages.
	imageData, imageType, err := image.Decode(f)
	if err != nil {
		return nil, "", fmt.Errorf("error decoding image: %w", err)
	}

	return imageData, imageType, nil
}

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	//	"image/gif"

	"math"
)

func main() {
	//height := 500
	//width := 500

	// Create a canvas to be painted
	var palette = []color.Color{color.Black, color.RGBA{G: 255, A: 255}}
	const (
		blackIndex = 0 // first color in palette
		greenIndex = 1 // next color in palette
	)

	const size = 1000

	canvas := image.NewPaletted(image.Rect(0, 0, size, size), palette)

	// Get a green crayon to paint
	//greenCrayon := color.RGBA{0, 255, 0, 200}

	for t := 0.0; t <= 12*math.Pi; t += 0.0001 {
		x := math.Sin(t) * (math.Pow(math.E, math.Cos(t)) - 2*math.Cos(4*t) - math.Pow(math.Sin(t/12), 5))
		y := math.Cos(t) * (math.Pow(math.E, math.Cos(t)) - 2*math.Cos(4*t) - math.Pow(math.Sin(t/12), 5))
		//fmt.Println(x, y)
		canvas.SetColorIndex(500+int((size/10)*x+0.5), 500+int((size/10)*y+0.5), greenIndex)
	}

	outputFile, _ := os.Create("butterfly.png")
	png.Encode(outputFile, canvas)
}

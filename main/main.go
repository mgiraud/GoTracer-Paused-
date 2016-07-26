package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"github.com/mgiraud/raytracer"
)

func main() {
	camera := &raytracer.Camera{}
	camera.SetPerspective(90, 0.01, 100, 16/9)
	camera.SetProjectionMatrix()
	fmt.Println(camera.ProjMat)
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))

	for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, color.RGBA{0x88, 0xff, 0x88, 0xff})
		}
	}

	file, err := os.Create("simple.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jpeg.Encode(file, img, &jpeg.Options{80})
}

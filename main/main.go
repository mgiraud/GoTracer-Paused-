package main

import (
	"path/filepath"

	"github.com/mgiraud/raytracer"
	"github.com/mgiraud/raytracer/matrix"
)

const IMG_WIDTH = 1920.0
const IMG_HEIGHT = 1080.0

func main() {
	path, _ := filepath.Abs("../src/github.com/mgiraud/raytracer/scenes/scene0.json")

	// Read Data (Camera, Shapes, Light) from provded json file
	camera, objects := raytracer.ReadScene(path)
	camera.InitCamera(IMG_WIDTH, IMG_HEIGHT, 100)
	camera.MoveCamera(matrix.Vector4{0, 0, 0, 1})

	// Creates the scene
	scene := &raytracer.Scene{
		Cam:     camera,
		Objects: map[string]raytracer.Object{},
	}
	scene.Objects = objects

	//Setup Objects Matrix
	for _, v := range scene.Objects {
		v.ComputeMatrix()
	}

	//Scene Rendering
	scene.Render(IMG_WIDTH, IMG_HEIGHT)

	//Creates img file
	scene.CreateImage()
}

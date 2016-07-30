package main

import (
	"path/filepath"

	"github.com/mgiraud/raytracer"
)

func main() {
	path, _ := filepath.Abs("../src/github.com/mgiraud/raytracer/scenes/scene0.json")

	// Read Data (Camera, Shapes, Light) from provded json file
	camera, objects, lights := raytracer.ReadScene(path)
	camera.InitCamera()

	// Creates the scene
	scene := &raytracer.Scene{}
	scene.InitScene(objects, lights, camera)

	//Setup Objects Matrix
	for _, v := range scene.Objects {
		v.ComputeMatrix()
	}

	//Setup Lights Matrix
	for _, l := range scene.Lights {
		l.InitLight()
	}

	//Scene Rendering
	scene.Render()

	//Creates img file
	scene.CreateImage()
}

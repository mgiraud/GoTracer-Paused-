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
	camera, objects := raytracer.ReadScene(path)
	camera.InitCamera(IMG_WIDTH, IMG_HEIGHT, 100)
	camera.MoveCamera(matrix.Vector4{0, 0, 0, 1})
	scene := &raytracer.Scene{Cam: camera}
	scene.Objects = map[string]raytracer.Object{}
	scene.Objects = objects

	sphereStruct := scene.Objects["sphere-1"]
	sphere := sphereStruct.(*raytracer.Sphere)
	sphere.ComputeMatrix()

	scene.Render(IMG_WIDTH, IMG_HEIGHT)
	scene.CreateImage()
}

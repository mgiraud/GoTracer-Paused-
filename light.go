package raytracer

import (
	"math"

	"github.com/mgiraud/raytracer/matrix"
)

type Light interface {
	InitLight()
}

type DefaultLight struct {
	Color        matrix.Vector4 `json:"color"`
	Intensity    float64        `json:"intensity"`
	Position     matrix.Vector4 `json:"position"`
	Rotation     matrix.Vector3 `json:"rotation"`
	lightToWorld matrix.Matrix16
	Dir          matrix.Vector4
	Pos          matrix.Vector4
}

type DistantLight struct {
	DefaultLight
}

type PointLight struct {
	DefaultLight
}

func (light *DistantLight) InitLight() {
	light.lightToWorld = matrix.NewMat16()
	light.lightToWorld.Mat16Identity()
	if light.Rotation[0] != 0 {
		light.lightToWorld.Rotate(light.lightToWorld, matrix.Vector3{1, 0, 0}, light.Rotation[0]*math.Pi/180)
	}
	if light.Rotation[1] != 0 {
		light.lightToWorld.Rotate(light.lightToWorld, matrix.Vector3{0, 1, 0}, light.Rotation[1]*math.Pi/180)
	}
	if light.Rotation[2] != 0 {
		light.lightToWorld.Rotate(light.lightToWorld, matrix.Vector3{0, 0, 1}, light.Rotation[2]*math.Pi/180)
	}
	light.Dir = matrix.NewVec4()
	light.Dir.Mat16MulVec4(light.lightToWorld, matrix.Vector4{0, 0, 1, 0})
}

func (light *PointLight) InitLight() {
	light.lightToWorld = matrix.NewMat16()
	light.lightToWorld.Mat16Identity()
	light.lightToWorld.Translate(light.lightToWorld, light.Position)
}

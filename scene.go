package raytracer

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"sync"

	"github.com/mgiraud/raytracer/matrix"
)

type Scene struct {
	Cam           *Camera
	img           *image.RGBA
	width, height float64
	Objects       map[string]Object
	mux           sync.Mutex
}

func (sce *Scene) GetIntersection(ray *Ray) (float64, Object) {
	d := math.MaxFloat64
	var o Object

	for _, v := range sce.Objects {
		o1 := v
		res, d1 := o1.Intersect(ray)
		if res == true && d1 < d {

			d = d1
			o = v
		}
	}
	return d, o
}

func (sce *Scene) CastRay(ray *Ray) color.RGBA {
	var col color.RGBA = color.RGBA{0, 0, 0, 255}

	d, o := sce.GetIntersection(ray)
	if d != math.MaxInt64 && o != nil {
		col = o.GetColor()
	}
	return col
}

func (sce *Scene) InitRay(i, j int) color.RGBA {
	cam := sce.Cam
	x := (2*(float64(i)+0.5)/sce.width - 1) * cam.ratio * cam.scale
	y := (1 - 2*(float64(j)+0.5)/sce.height) * cam.scale
	dir := matrix.Vector4{0, 0, 0, 0}
	dir.Mat16MulVec4(cam.CamToWorld, matrix.Vector4{x, y, 1, 0})
	dir.Normalize()
	ori := matrix.Vector4{0, 0, 0, 1}
	ori.Mat16MulVec4(cam.CamToWorld, ori)
	ray := &Ray{
		I:   i,
		J:   j,
		Ori: ori,
		Dir: dir,
	}
	return sce.CastRay(ray)
}

func (sce *Scene) RenderBlock(x int, y int) {

	for z := y; z <= y+8; z++ {
		for w := x; w <= x+8; w++ {
			col := sce.InitRay(w, z)
			sce.img.Set(w, z, col)
		}
	}

}

func (sce *Scene) Render(width, height float64) {
	sce.img = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	sce.width = width
	sce.height = height
	var wg sync.WaitGroup
	wg.Add(sce.img.Rect.Max.Y * sce.img.Rect.Max.X / 64)

	for y := sce.img.Rect.Min.Y; y < sce.img.Rect.Max.Y; y += 8 {
		for x := sce.img.Rect.Min.X; x < sce.img.Rect.Max.X; x += 8 {
			a, b := x, y
			go func() {
				defer wg.Done()
				sce.RenderBlock(a, b)
			}()
		}
	}
	wg.Wait()
}

func (sce *Scene) CreateImage() {
	file, err := os.Create("simple.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jpeg.Encode(file, sce.img, &jpeg.Options{80})
}

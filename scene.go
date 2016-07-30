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
	Cam     *Camera
	img     *image.RGBA
	Objects map[string]Object
	Lights  map[string]Light
	mux     sync.Mutex
}

func (sce *Scene) IsIlluminatedByDirlight(light *DistantLight, i matrix.Vector4, n matrix.Vector4) bool {
	invDir := light.Dir.Neg()
	ray := NewRay(i.Add(n.MulFloat(sce.Cam.Bias)), invDir, "SHADOW_DIRECT_LIGHT")
	_, _, hasShadow := sce.GetIntersection(ray)
	return hasShadow == nil
}

func (sce *Scene) ComputeDirectLight(o Object, i matrix.Vector4, n matrix.Vector4) matrix.Vector4 {
	var diffuse matrix.Vector4 = matrix.Vector4{0, 0, 0, 255}
	var specular = matrix.Vector4{0, 0, 0, 255}
	var diffuseFactor float64
	var specularFactor float64
	for _, l := range sce.Lights {
		d, ok := l.(*DistantLight)
		if ok {
			invdir := d.Dir.Neg()
			vis := sce.IsIlluminatedByDirlight(d, i, n)
			if vis {
				diffuseFactor = o.GetAlbedo() / math.Pi * math.Max(0.0, n.Dot(invdir)) * d.Intensity
				R := Reflect(d.Dir, n)
				specularFactor = d.Intensity * math.Pow(math.Max(0.0, R.Dot(invdir)), o.GetN())
			} else {
				diffuseFactor = 0.0
				specularFactor = 0.0
			}
			diffuse[0] += d.Color[0] * o.GetColor()[0] * diffuseFactor
			diffuse[1] += d.Color[1] * o.GetColor()[1] * diffuseFactor
			diffuse[2] += d.Color[2] * o.GetColor()[2] * diffuseFactor

			specular[0] += d.Color[0] * specularFactor
			specular[1] += d.Color[1] * specularFactor
			specular[2] += d.Color[2] * specularFactor
		}
	}
	return matrix.Vector4{
		diffuse[0]*o.GetKd() + specular[0]*o.GetKs(),
		diffuse[1]*o.GetKd() + specular[1]*o.GetKs(),
		diffuse[2]*o.GetKd() + specular[2]*o.GetKs(),
		255,
	}
}

func (sce *Scene) GetIntersection(ray *Ray) (float64, Object, matrix.Vector4) {
	d := math.MaxFloat64
	var o Object = nil
	var i matrix.Vector4 = nil

	for _, v := range sce.Objects {
		o1 := v
		res, d1, i1 := o1.Intersect(ray)
		if res == true && d1 < d {
			i = i1
			d = d1
			o = v
		}
	}
	return d, o, i
}

func (sce *Scene) CastRay(ray *Ray) (bool, matrix.Vector4) {
	var col matrix.Vector4 = matrix.Vector4{0, 0, 0, 255}
	var hasIntersec bool = false

	d, o, i := sce.GetIntersection(ray)
	if d != math.MaxInt64 && o != nil {
		n := o.Normale(i, ray)
		col = sce.ComputeDirectLight(o, i, n)
		hasIntersec = true
	}
	return hasIntersec, col
}

func Vec4ToRGBA(col matrix.Vector4) color.RGBA {
	return color.RGBA{
		R: uint8(math.Min(255, col[0])),
		G: uint8(math.Min(255, col[1])),
		B: uint8(math.Min(255, col[2])),
		A: uint8(math.Min(255, col[3])),
	}
}

func (sce *Scene) InitRay(i, j int) {
	cam := sce.Cam
	x := (2*(float64(i)+0.5)/sce.Cam.Width - 1) * cam.ratio * cam.scale
	y := (1 - 2*(float64(j)+0.5)/sce.Cam.Height) * cam.scale
	dir := matrix.Vector4{0, 0, 0, 0}
	dir.Mat16MulVec4(cam.CamToWorld, matrix.Vector4{x, y, 1, 0})
	dir.Normalize()
	ori := matrix.Vector4{0, 0, 0, 1}
	ori.Mat16MulVec4(cam.CamToWorld, ori)
	ray := NewRay(ori, dir, "PRIMARY")
	_, col := sce.CastRay(ray)
	sce.mux.Lock()
	sce.img.Set(i, j, Vec4ToRGBA(col))
	sce.mux.Unlock()
}

func (sce *Scene) RenderBlock(x int, y int) {

	for z := y; z <= y+8; z++ {
		for w := x; w <= x+8; w++ {
			sce.InitRay(w, z)
		}
	}

}

func (sce *Scene) Render() {
	sce.img = image.NewRGBA(image.Rect(0, 0, int(sce.Cam.Width), int(sce.Cam.Height)))
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

func (sce *Scene) InitScene(objects ObjectMap, lights LightMap, cam *Camera) {
	sce.Cam = cam
	// sce.Objects = map[string]raytracer.Object{}
	// sce.Lights =  map[string]raytracer.Light{}
	sce.Objects = objects
	sce.Lights = lights
}

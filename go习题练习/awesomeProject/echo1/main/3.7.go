package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

type Func func(complex128) complex128

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, z4(z))
		}
	}
	png.Encode(os.Stdout, img)
	// 创建文件
	file, err := os.Create("3.7.png")
	if err != nil {
		log.Fatal(err)
	}

	// 使用png格式将数据写入文件
	png.Encode(file, img) //将image信息写入文件中

	// 关闭文件
	file.Close()
}

func z4(z complex128) color.Color {
	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}
	fPrime := func(z complex128) complex128 {
		return (z - 1/(z*z*z)) / 4
	}
	return newton(z, f, fPrime)
}

func newton(z complex128, f Func, fPrime Func) color.Color {
	const iterations = 200
	const contrast = 15
	for n := uint8(0); n < iterations; n++ {
		z -= fPrime(z)
		if cmplx.Abs(f(z)) < 1e-6 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}

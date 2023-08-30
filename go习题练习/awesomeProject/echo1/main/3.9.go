package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

// web服务器，用于给客户端生成分形的图像
func main() {
	http.HandleFunc("/", hander)
	if err := http.ListenAndServe(":8002", nil); err != nil {
		fmt.Println("server listen error,err:", err)
	}
}

func Test(f io.Writer, width int, height int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(f, img) // NOTE: ignoring errors
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
func hander(w http.ResponseWriter, r *http.Request) {
	// 从HTTP参数中获取长、宽、放大倍数
	r.ParseForm()
	width, height, zoom := 1024, 1024, 1
	if s := r.Form.Get("x"); s != "" {
		width, _ = strconv.Atoi(s)
	}
	if s := r.Form.Get("y"); s != "" {
		height, _ = strconv.Atoi(s)
	}
	if s := r.Form.Get("zoom"); s != "" {
		zoom, _ = strconv.Atoi(s)
	}
	Test(w, width*zoom, height*zoom)
}

// Surface computes an SVG redenring of a 3-D surface function.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

func main() {

	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, width, height))

	//遍历每一个元素
	for x := 0; x < width; x++ {

		for y := 0; y < height; y++ {
			//填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; stroke-width: 0.7' "+"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			pic.SetGray(int(ax), int(ay), color.Gray{0})
			pic.SetGray(int(bx), int(by), color.Gray{0})
			pic.SetGray(int(cx), int(cy), color.Gray{0})
			pic.SetGray(int(dx), int(dy), color.Gray{0})
		}
	}

	// 创建文件
	file, err := os.Create("3.2_eggbox.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()

}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*math.Cos(angle)*xyscale
	sy := height/2 + (x+y)*math.Sin(angle)*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := 0.2 * (math.Cos(x) + math.Cos(y))
	return r
}

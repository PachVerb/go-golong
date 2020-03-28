package main

import (
	"image"
	"image/color" //image颜色包
	"image/png"   //解码图像数据
	"math/cmplx"
	"os"
)

func main() {
	//var out io.Writer
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height)) //返回具有给定边界的图像

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	//png.Encode(os.Stdout, img) // NOTE: ignoring errors, 这里得编码图像输出到了控制台，现在重定向到io.writer
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.NRGBA{R: 255 - contrast*n, G: 1, B: 2, A: 4}
		}
	}
	return color.RGBA{R: 66, G: 144, B: 255, A: 1}
}

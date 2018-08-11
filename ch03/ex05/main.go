package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

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
			// 画像の点 (px, py) は複素数値z を表している
			if len(os.Args) > 1 && os.Args[1] == "ycbcr" {
				img.Set(px, py, mandelbrotYCbCr(z))
			} else {
				img.Set(px, py, mandelbrotRGBA(z))
			}
		}
	}
	png.Encode(os.Stdout, img) // 注意: エラーを無視
}

func mandelbrotRGBA(z complex128) color.Color {
	const iterations = 200
	const contrastR = 25
	const contrastG = 40
	const contrastB = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			red := uint8(255 - contrastR*n)
			green := uint8(255 - contrastG*n)
			blue := uint8(255 - contrastB*n)
			return color.RGBA{red, green, blue, 255}
		}
	}
	return color.Black
}

func mandelbrotYCbCr(z complex128) color.Color {
	const iterations = 200
	const contrastY = 50
	const contrastCb = 15
	const contrastCr = 100

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			red := uint8(255 - contrastY*n)
			green := uint8(255 - contrastCb*n)
			blue := uint8(255 - contrastCr*n)
			return color.YCbCr{red, green, blue}
		}
	}
	return color.Black
}

package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"os"
)

WIP

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img64 := image.NewRGBA(image.Rect(0, 0, width, height))
	img128 := image.NewRGBA(image.Rect(0, 0, width, height))
	// imgfloat := image.NewRGBA(image.Rect(0, 0, width, height))
	// imgrat := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y128 := float64(py)/height*(ymax-ymin) + ymin
		y64 := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x128 := float64(px)/width*(xmax-xmin) + xmin
			x64 := float32(px)/width*(xmax-xmin) + xmin
			z128 := complex(x128, y128)
			z64 := complex(x64, y64)
			// 画像の点 (px, py) は複素数値z を表している
			img64.Set(px, py, mandelbrot64(z64))
			img128.Set(px, py, mandelbrot128(z128))
		}
	}
	writePng("out64.png", img64)
	writePng("out128.png", img128)
}

func mandelbrot128(z complex128) color.Color {
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

func mandelbrot64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if math.Hypot(float64(real(v)), float64(imag(v))) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

/*
func mandelbrotBigFloat(z big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	var v big.Float
	for n := uint8(0); n < iterations; n++ {
		v = v.Mul(v).Add()
		if v > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotBigRat(z big.Rat) color.Color {
	const iterations = 200
	const contrast = 15

	var v big.Rat
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
*/

func writePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := io.Writer(file)
	png.Encode(writer, img)
}

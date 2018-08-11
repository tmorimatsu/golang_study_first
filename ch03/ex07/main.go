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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // 注意: エラーを無視
}

func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	const threshold = 1e-6

	for i := uint8(0); i < iterations; i++ {
		z = z - (z*z*z*z-1)/(4*z*z*z)
		if cmplx.Abs(z-complex(1, 0)) < threshold {
			return color.RGBA{uint8(255 - i), 0, 0, 255}
		}
		if cmplx.Abs(z-complex(-1, 0)) < threshold {
			return color.RGBA{0, uint8(255 - i), 0, 255}
		}
		if cmplx.Abs(z-complex(0, 1)) < threshold {
			return color.RGBA{0, 0, uint8(255 - i), 255}
		}
		if cmplx.Abs(z-complex(0, -1)) < threshold {
			return color.RGBA{uint8(i), uint8(i), uint8(i), 255}
		}
	}
	return color.Black
}

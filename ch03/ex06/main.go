package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

/*
* スーパーサンプリングは、ここの画素内の複数の点のカラー地を計算して平均を求めることでピクセル化の影響を薄める技法です。
* もっとも単純な方法は、ここの画素を四つの「サブピクセル」へ分割することです。その方法を実装しなさい。
 */

// 問題の意味があまりよくわからない

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	// 元の画像の要素の退避先を作成
	c := make([][]color.RGBA, width)
	for i := range c {
		c[i] = make([]color.RGBA, height)
	}

	// 元の画像の要素を生成
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値z を表している
			c[px][py] = mandelbrotRGBA(z)
		}
	}

	// 加工した要素をset
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			img.Set(px, py, superSampling(c, px, py))
		}
	}

	png.Encode(os.Stdout, img) // 注意: エラーを無視
}

func superSampling(c [][]color.RGBA, x int, y int) color.Color {

	// 周りに4つ要素がない箇所はそのままの値を返却
	if x == 0 || y == 0 || x == len(c)-1 || y == len(c[0])-1 {
		return c[x][y]
	}

	r := (int(c[x+1][y].R) + int(c[x-1][y].R) + int(c[x][y+1].R) + int(c[x][y-1].R)) / 4
	g := (int(c[x+1][y].G) + int(c[x-1][y].G) + int(c[x][y+1].G) + int(c[x][y-1].G)) / 4
	b := (int(c[x+1][y].B) + int(c[x-1][y].B) + int(c[x][y+1].B) + int(c[x][y-1].B)) / 4

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func mandelbrotRGBA(z complex128) color.RGBA {
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
	return color.RGBA{0, 0, 0, 255}
}

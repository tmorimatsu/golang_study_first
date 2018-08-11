package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	// パラメーターの構造体をクエリから取得し、メソッドに渡す
	zoom := getParam(r.Form, "zoom", 2)
	xmin := getParam(r.Form, "x", 0) - zoom
	ymin := getParam(r.Form, "y", 0) - zoom
	xmax := getParam(r.Form, "x", 0) + zoom
	ymax := getParam(r.Form, "y", 0) + zoom

	const (
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値z を表している
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // 注意: エラーを無視
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

func getParam(v url.Values, key string, defaultValue float64) float64 {

	value, err := strconv.ParseFloat(v.Get(key), 64)
	if err != nil {
		log.Print(key)
		log.Print(err)
		return defaultValue
	}
	return value
}

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
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
	type Params struct {
		cycles  int     // number of complete x oscillator revolutions
		res     float64 // angular resolution
		size    int     // image canvas covers [-size..+size]
		nframes int     // number of animation frames
		delay   int     // delay between frames in 10ms units
	}
	var test Params
	var err error
	test.cycles, err = strconv.Atoi(r.Form.Get("cycles"))
	if err != nil {
		log.Print(err)
		test.cycles = 5
	}
	lissajous(w, float64(test.cycles))
}

var palette = []color.Color{color.Black, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

const (
	blackIndex = 0 // パレットの最初の色
	greenIndex = 1 // パレットの次の色
	blueIndex  = 2 // パレットの次の色
)

func lissajous(out io.Writer /*params *test*/, cycles float64) {
	const (
		//cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.0001 // angular resolution
		size    = 100    // image canvas covers [-size..+size]
		nframes = 64     // number of animation frames
		delay   = 8      // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

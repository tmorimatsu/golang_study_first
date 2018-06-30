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
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type Params struct {
	cycles  float64 // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	// パラメーターの構造体をクエリから取得し、メソッドに渡す
	var params Params
	params.cycles = getMapData(r.Form, "cycles", 5.0)
	params.res = getMapData(r.Form, "res", 0.0001)
	params.size = int(getMapData(r.Form, "size", 100))
	params.nframes = int(getMapData(r.Form, "nframes", 64))
	params.delay = int(getMapData(r.Form, "delay", 8))

	lissajous(w, params)
}

func getMapData(m url.Values, key string, defaultValue float64) float64 {

	// TODO: 値に上限値, 下限値を決めてバリデーション
	// 処理に時間がかかりすぎて終わらないようになる可能性があるため

	value, err := strconv.ParseFloat(m.Get(key), 64)
	if err != nil {
		log.Print(key)
		log.Print(err)
		return defaultValue
	}
	return value
}

var palette = []color.Color{color.Black, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

const (
	blackIndex = 0 // パレットの最初の色
	greenIndex = 1 // パレットの次の色
	blueIndex  = 2 // パレットの次の色
)

func lissajous(out io.Writer, params Params) {

	cycles, res, size, nframes, delay := params.cycles, params.res, params.size, params.nframes, params.delay

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

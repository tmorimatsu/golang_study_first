package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type Params struct {
	width   float64
	height  float64
	cells   float64
	xyrange float64
	xyscale float64
	zscale  float64
	angle   float64
}

var params Params

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	params.width = getValue(r.Form, "width", 600)
	params.height = getValue(r.Form, "height", 320)
	setParams()

	writeSvg(w)
}

func writeSvg(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d' >", params.width, params.height)
	for i := 0; i < int(params.cells); i++ {
		for j := 0; j < int(params.cells); j++ {
			ax, ay, acolor, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, _, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, _, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, _, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g, %g %g, %g %g,%g %g, %g' fill='%s' /> \n", ax, ay, bx, by, cx, cy, dx, dy, acolor)
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, string, error) {
	var sin30, cos30 = math.Sin(params.angle), math.Cos(params.angle)

	x := params.xyrange * (float64(i)/params.cells - 0.5)
	y := params.xyrange * (float64(j)/params.cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, "", fmt.Errorf("infinity")
	}

	if math.IsNaN(z) {
		return 0, 0, "", fmt.Errorf("not a number")
	}

	sx := params.width/2 + (x-y)*cos30*params.xyscale
	sy := params.height/2 + (x+y)*sin30*params.xyscale - z*params.zscale
	return sx, sy, color(z), nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func color(z float64) string {
	if z < 0 {
		return "#0000FF"
	}
	v := int(z * 255)
	return fmt.Sprintf("#%02x00%02x", v, 255-v)
}

func setParams() {

	params.cells = 100
	params.xyrange = 30.0
	params.xyscale = params.width / 2.0 / params.xyrange
	params.zscale = params.height * 0.4
	params.angle = math.Pi / 6
}

func getValue(m url.Values, key string, defaultValue float64) float64 {

	value, err := strconv.ParseFloat(m.Get(key), 64)
	if err != nil {
		log.Print(key)
		log.Print(err)
		return defaultValue
	}
	return value
}

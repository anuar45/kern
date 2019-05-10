package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

type svgParams struct {
	xyrange float64
	xyscale float64
	zscale  float64
	height  float64
	width   float64
	cells   int
	angle   float64
	sin30   float64
	cos30   float64
}

// var xyrange float 64
// var xyscale, zscale, height float64
// var height, width int
// cells = 100
// angle = math.Pi / 6
// var sin30, cos30 = math.Sin(angle), math.Cos(angle)
// xyrange = 30.0
// xyscale = float64(width) / 2.0 / xyrange
// zscale = float64(height) * 0.4

func main() {

	http.HandleFunc("/surface", handler)
	http.ListenAndServe("localhost:8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	height, _ := strconv.ParseFloat(r.Form["h"][0], 64)
	width, _ := strconv.ParseFloat(r.Form["w"][0], 64)

	s := newSvgParams(height, width)

	svg := getSvg(s)
	fmt.Fprintf(w, svg)
}

func newSvgParams(height, width float64) svgParams {
	xyrange := 30.0
	xyscale := width / 2.0 / xyrange
	zscale := height * 0.4
	angle := math.Pi / 6
	sin30 := math.Sin(angle)
	cos30 := math.Cos(angle)

	s := svgParams{
		xyrange: 30.0,
		xyscale: xyscale,
		zscale:  zscale,
		height:  height,
		width:   width,
		cells:   100,
		angle:   math.Pi / 6,
		sin30:   sin30,
		cos30:   cos30,
	}

	return s
}

func getSvg(s svgParams) string {

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", s.width, s.height)

	for i := 0; i < s.cells; i++ {
		for j := 0; j < s.cells; j++ {
			ax, ay := corner(s, i+1, j)
			bx, by := corner(s, i, j)
			cx, cy := corner(s, i, j+1)
			dx, dy := corner(s, i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g '/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(s svgParams, i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i, j).
	x := s.xyrange * (float64(i)/float64(s.cells) - 0.5)
	y := s.xyrange * (float64(j)/float64(s.cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := s.width/2 + (x-y)*s.cos30*s.xyscale
	sy := s.height/2 + (x+y)*s.sin30*s.xyscale - z*s.zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

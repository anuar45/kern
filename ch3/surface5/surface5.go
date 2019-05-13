package main

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

type svgParams struct {
	xyrange float64
	xyscale float64
	zscale  float64
	height  int
	width   int
	cells   int
	angle   float64
	sin30   float64
	cos30   float64
}

// http://localhost:8080/surface?h=400&w=600

func main() {

	http.HandleFunc("/surface", handler)
	http.ListenAndServe("localhost:8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	height, _ := strconv.Atoi(r.Form["h"][0])
	width, _ := strconv.Atoi(r.Form["w"][0])

	s := newSvgParams(height, width)

	svg := getSvg(s)
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, svg)
}

func newSvgParams(height, width int) svgParams {
	xyrange := 40.0
	xyscale := float64(width) / 2.0 / xyrange
	zscale := float64(height) * 0.05
	angle := math.Pi / 6
	sin30 := math.Sin(angle)
	cos30 := math.Cos(angle)

	s := svgParams{
		xyrange: xyrange,
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
	var svg bytes.Buffer

	svg.WriteString(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", s.width, s.height))

	for i := 0; i < s.cells; i++ {
		for j := 0; j < s.cells; j++ {
			ax, ay, color := corner(s, i+1, j)
			bx, by, _ := corner(s, i, j)
			cx, cy, _ := corner(s, i, j+1)
			dx, dy, _ := corner(s, i+1, j+1)
			svg.WriteString(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color))
		}
	}

	svg.WriteString(fmt.Sprintln("</svg>"))

	return svg.String()
}

func corner(s svgParams, i, j int) (float64, float64, string) {
	// Find point (x,y) at corner of cell (i, j).
	x := s.xyrange * (float64(i)/float64(s.cells) - 0.5)
	y := s.xyrange * (float64(j)/float64(s.cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := float64(s.width)/2.0 + (x-y)*s.cos30*s.xyscale
	sy := float64(s.height)/2.0 + (x+y)*s.sin30*s.xyscale - z*s.zscale

	color := "blue"

	if z >= 0 {
		color = "red"
	}

	return sx, sy, color
}

func f(x, y float64) float64 {
	// r := math.Hypot(x, y) // distance from (0,0)
	// r := x + y // distance from (0,0)
	// return math.Sin(r) / r
	return math.Sin(x) + math.Sin(y)
}

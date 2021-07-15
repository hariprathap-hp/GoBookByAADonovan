package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	cells         = 200
	angle         = math.Pi / 6
	xyrange       = 30
	width, height = 600, 320
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {

	http.HandleFunc("/svg/", svgHandler)
	log.Fatalln(http.ListenAndServe(":8000", nil))

}

func svgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, fmt.Errorf("invalid value")
	}
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 16
}

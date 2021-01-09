package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			if math.IsNaN(ax) || math.IsNaN(ay) {
				continue
			}
			bx, by := corner(i, j)
			if math.IsNaN(bx) || math.IsNaN(by) {
				continue
			}
			cx, cy := corner(i, j+1)
			if math.IsNaN(cx) || math.IsNaN(cy) {
				continue
			}
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color(i, j))
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func color(i, j int) string {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y) * 10
	if z > 0 {
		num := int(z*255) % 255
		return fmt.Sprintf("#00%x00", num)
	} else {
		num := int(-z*255) % 255
		return fmt.Sprintf("#%x0000", num)
	}
}

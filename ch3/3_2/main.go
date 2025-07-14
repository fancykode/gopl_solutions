/* Exercise 3.2: Experiment with visualizations of other functions from the math package. Can
   you produce an eggbox, moguls, or a saddle?
*/
package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	width, height = 600, 320            //canvas size in pixels
	cells         = 100                 // nuber of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

var mogulsFlag = flag.Bool("moguls", false, "moguls")
var eggboxFlag = flag.Bool("eggbox", false, "eggbox")
var saddleFlag = flag.Bool("saddle", false, "saddle")

func main() {
	flag.Parse()
	if *mogulsFlag {
		draw(moguls)
	} else if *eggboxFlag {
		draw(eggbox)
	} else if *saddleFlag {
		draw(saddle)
	}
}

func draw(f func(float64, float64) float64) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func isFinite(n float64) bool {
	return !math.IsNaN(n) && !math.IsInf(n, 0)
}

func corner(i, j int, f func(float64, float64) float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func saddle(x, y float64) float64 {
	return (x*x - y*y) / 500
}

func eggbox(x, y float64) float64 {
	return (math.Sin(x/2.0) + math.Sin(y/2.0)) * 0.3
}

func moguls(x, y float64) float64 {
	return math.Cos(math.Sqrt(x*x+y*y)) * 0.1
}

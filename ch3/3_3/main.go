/* Exercise 3.3: Color each polygon based on its height, so that the peaks are colored red
(#ff0000) and the valleys blue (#0000ff).
*/
package main

import (
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

func calcMinMaxHeight() (float64, float64) {
	minH := math.NaN()
	maxH := math.NaN()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z := calcXYZ(i, j)
			if isFinite(z) {
				if math.IsNaN(minH) || z < minH {
					minH = z
				}
				if math.IsNaN(maxH) || z > maxH {
					maxH = z
				}
			}
		}
	}
	return minH, maxH
}

func main() {

	minHeight, maxHeight := calcMinMaxHeight()

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, h := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				r, b := calcColor(h, minHeight, maxHeight)
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"rgb(%d, 0, %d)\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, r, b)
			}
		}
	}
	fmt.Println("</svg>")
}

func isFinite(n float64) bool {
	return !math.IsNaN(n) && !math.IsInf(n, 0)
}

// returns red ang blue colors
func calcColor(h, minHeight, maxHeight float64) (int, int) {
	if !isFinite(h) || !isFinite(minHeight) || !isFinite(maxHeight) {
		return 0, 0
	}
	val := int(((h - minHeight) / (maxHeight - minHeight)) * 255)
	red := val
	blue := 255 - val
	return red, blue
}

func calcXYZ(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	return x, y, z
}

func corner(i, j int) (float64, float64, float64) {
	x, y, z := calcXYZ(i, j)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

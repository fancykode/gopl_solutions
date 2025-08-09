/* Exercise 3.4: Following the approach of the Lissajous example in Section 1.7, construct a web
   server that computes surfaces and writes SVG data to the client. The server must set
   the Content-Type he ader like this: w.Header().Set("Content-Type", "image/svg+xml")
   (This step was not required in the Lissajous example because the server uses standard
   heuristics to recognize common formats like PNG from the first 512 bytes of the response, and
   generates the proper header.) Allow the client to specify values like height, width, and color
   as HTTP request parameters.
*/
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // nuber of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30)

)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		heightParam := r.URL.Query().Get("height")
		widthParam := r.URL.Query().Get("width")
		colorParam := r.URL.Query().Get("color")
		color := "grey"
		if colorParam != "" {
			color = colorParam
		}

		height, err := strconv.Atoi(heightParam)
		if err != nil {
			height = 320
		}
		width, err := strconv.Atoi(widthParam)
		if err != nil {
			width = 600
		}

		xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
		zscale := float64(height) * 0.4         // pixels per z unit
		w.Header().Set("Content-Type", "image/svg+xml")
		genSurface(w, float64(width), float64(height), xyscale, zscale, color)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func genSurface(out io.Writer, width, height, xyscale, zscale float64, color string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, xyscale, zscale)
			bx, by := corner(i, j, width, height, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, xyscale, zscale)
			if isFinite(ax) && isFinite(ay) &&
				isFinite(bx) && isFinite(by) &&
				isFinite(cx) && isFinite(cy) &&
				isFinite(dx) && isFinite(dy) {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func isFinite(n float64) bool {
	return !math.IsNaN(n) && !math.IsInf(n, 0)
}

func corner(i, j int, width, height, xyscale, zscale float64) (float64, float64) {
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

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	result := math.Sin(r) / r
	return result
}

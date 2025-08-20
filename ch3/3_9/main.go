/* Exercise 3.9: Write a webserver that renders fractals and writes the image data to the client.
   Allow the client to specify the x, y, and zoom values as parameters to the HTTP request.
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 255},
	color.RGBA{0x00, 0x70, 0xff, 255},
	color.RGBA{0x00, 0x00, 0xff, 255},
	color.RGBA{0xdc, 0x14, 0x3c, 255},
	color.RGBA{0xff, 0x8c, 0x00, 255},
	color.RGBA{0x00, 0x80, 0x00, 255},
	color.RGBA{0x80, 0x00, 0x80, 255},
	color.RGBA{0xff, 0x00, 0x00, 255},
	color.RGBA{0xff, 0xff, 0x00, 255},
	color.RGBA{0x00, 0x00, 0x80, 255},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cxParam := r.URL.Query().Get("x")
		cyParam := r.URL.Query().Get("y")
		zoomParam := r.URL.Query().Get("zoom")
		cx, err := strconv.ParseFloat(cxParam, 64)
		if err != nil {
			cx = 0
		}

		cy, err := strconv.ParseFloat(cyParam, 64)
		if err != nil {
			cy = 0
		}
		zoom, err := strconv.ParseFloat(zoomParam, 64)
		if err != nil {
			zoom = 1
		}
		renderMandelbrot(w, cx, cy, zoom)
	})
	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func renderMandelbrot(out io.Writer, cx, cy, zoom float64) {
	const (
		width, height = 1024, 1024
	)
	var xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
	dx := (xmax - xmin) / zoom
	dy := (ymax - ymin) / zoom
	xmin = cx - dx/2.0
	xmax = cx + dx/2.0
	ymin = cy - dy/2.0
	ymax = cy + dy/2.0

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	var v complex128
	for i := 0; i < iterations; i++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[i%len(palette)]
		}
	}
	return color.Black
}

/* Exercise 3.6: Supersampling is a technique to reduce the effect of pixelation by computing the
color value at several points within each pixel and taking the average. The simplest method is
to divide each pixel into four "subpixels." Implement it.
*/
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := (float64(px)+0.5)/width*(xmax-xmin) + xmin
			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)
			c := avrg([]color.Color{
				mandelbrot(z1),
				mandelbrot(z2),
				mandelbrot(z3),
				mandelbrot(z4),
			})
			// Image point (px, py) represents complex value z.
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func avrg(colors []color.Color) color.Color {
	var r, g, b, a float64
	if len(colors) == 0 {
		return nil
	}

	for _, c := range colors {
		tr, tg, tb, ta := c.RGBA()
		r += float64(tr>>8) / float64(len(colors))
		g += float64(tg>>8) / float64(len(colors))
		b += float64(tb>>8) / float64(len(colors))
		a += float64(ta>>8) / float64(len(colors))
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 255 - contrast*n, 0, 255}
		}
	}
	return color.Black
}

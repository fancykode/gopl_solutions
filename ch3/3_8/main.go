/* Exercise 3.8: Rendering fractals at high zoom levels demands great arithmetic precision.
   Implement the same fractal using four different represent ations of numbers: complex64,
   complex128, big.Float, and big.Rat. (The latter two types are found in the math/big package.
   Float uses arbitrary but bounded-precision floating-point; Rat uses unbounded-precision
   rational numbers.) How do they compare in performance and memory usage? At what zoom
   levels do rendering artifacts become visible?
*/
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

var m64 = flag.Bool("m64", false, "Mandelbrot computed with complex64 type")
var m128 = flag.Bool("m128", false, "Mandelbrot computed with complex128 type")
var mBigFloat = flag.Bool("mbf", false, "Mandelbrot computed with big.Float type")
var mBigRat = flag.Bool("mbr", false, "Mandelbrot computed with big.Rat type")

func main() {
	flag.Parse()
	if *m64 {
		calcMandelbrot64(os.Stdout)
	} else if *m128 {
		calcMandelbrot128(os.Stdout)
	} else if *mBigFloat {
		calcMandelbrotBigFloat(os.Stdout)
	} else if *mBigRat {
		calcMandelbrotBigRat(os.Stdout)
	}
}

func calcMandelbrot64(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			var z complex64 = complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot64(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func calcMandelbrot128(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot128(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func calcMandelbrotBigFloat(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotBigFloat(x, y))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func calcMandelbrotBigRat(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotBigRat(x, y))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrotBigFloat(real, imag float64) color.Color {
	const iterations = 200
	const contrast = 15
	zRe := big.NewFloat(real)
	zIm := big.NewFloat(imag)
	vRe := big.NewFloat(0.0)
	vIm := big.NewFloat(0.0)
	for n := uint8(0); n < iterations; n++ {
		tvRe, tvIm := &big.Float{}, &big.Float{}
		tvRe.Mul(vRe, vRe).Sub(tvRe, (&big.Float{}).Mul(vIm, vIm)).Add(tvRe, zRe)
		tvIm.Mul(vRe, vIm).Mul(tvIm, big.NewFloat(2.0)).Add(tvIm, zIm)
		vRe, vIm = tvRe, tvIm
		absVal := &big.Float{}
		absVal.Mul(vRe, vRe).Add(absVal, (&big.Float{}).Mul(vIm, vIm)).Sqrt(absVal)
		if absVal.Cmp(big.NewFloat(2)) == 1 {
			return color.RGBA{0, 255 - contrast*n, 0, 255}
		}

	}
	return color.Black
}

func mandelbrotBigRat(real, imag float64) color.Color {
	const iterations = 200
	const contrast = 15
	zRe := (&big.Rat{}).SetFloat64(real)
	zIm := (&big.Rat{}).SetFloat64(imag)
	vRe := &big.Rat{}
	vIm := &big.Rat{}
	for n := uint8(0); n < iterations; n++ {
		tvRe, tvIm := &big.Rat{}, &big.Rat{}
		tvRe.Mul(vRe, vRe).Sub(tvRe, (&big.Rat{}).Mul(vIm, vIm)).Add(tvRe, zRe)
		tvIm.Mul(vRe, vIm).Mul(tvIm, big.NewRat(2, 1)).Add(tvIm, zIm)
		vRe, vIm = tvRe, tvIm
		absVal := &big.Rat{}
		absVal.Mul(vRe, vRe).Add(absVal, (&big.Rat{}).Mul(vIm, vIm))
		if absVal.Cmp(big.NewRat(4, 1)) == 1 {
			return color.RGBA{0, 255 - contrast*n, 0, 255}
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
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

func mandelbrot64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.RGBA{0, 255 - contrast*n, 0, 255}
		}
	}
	return color.Black
}

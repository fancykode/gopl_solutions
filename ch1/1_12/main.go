/*
Exercise 1.12: Modify the Lissajous server to read parameter values from the URL.
For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the
number of cycles to 20 instead of the default 5. Use the strconv.Atoi function to convert the
string parameter into an integer. You can see its documentation with go doc strconv.Atoi.
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cyclesParam := r.URL.Query().Get("cycles")
		sizeParam := r.URL.Query().Get("size")
		cycles, err := strconv.Atoi(cyclesParam)
		if err != nil {
			cycles = 5
		}
		size, err := strconv.Atoi(sizeParam)
		if err != nil {
			size = 100
		}
		lissajous(w, cycles, size)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int, size int) {
	const (
		res     = 0.001 // angular resolution
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

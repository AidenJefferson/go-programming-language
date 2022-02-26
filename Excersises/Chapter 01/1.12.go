// This program starts a server on localhost:8000 which produces a lissajous diagram, the diagram can be changed using 
// the queries ?cycles=50&res=0.001&size=400&nframes=24&delay=1&freq=0.765 (not all queries are required)
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	
 	lissa_vars := make(map[string]float64)

	// declare default variables
	lissa_vars["cycles"] = 8
	lissa_vars["res"] = 0.001
	lissa_vars["size"] = 300
	lissa_vars["nframes"] = 64
	lissa_vars["delay"] = 1
	lissa_vars["freq"] = 4

	input_lissa := r.URL.Query()

	// // update variables with values given
	for key, value := range input_lissa{
	 	lissa_vars[key], _ = strconv.ParseFloat(value[0], 64) // only allow first entry to be used
	}

	mu.Lock()
	count++
	lissajous(w, lissa_vars["cycles"], lissa_vars["res"], lissa_vars["size"], lissa_vars["nframes"], lissa_vars["delay"],  lissa_vars["freq"])
	mu.Unlock()
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// creates a lissajous diagram
	// cycles - number of complete x oscillator revolutions
	// res - angular resolution
	// size - image canvas covers [-size..+size]
	// nframes - number of animation frames
	// delay - delay between frames in 10ms units
func lissajous(out io.Writer, cycles float64, res float64, size float64, nframes float64, delay float64, freq float64) {
	var palette = []color.Color{color.Black, color.RGBA{0, 128, 0, 1}}
	const (
		greenIndex = 0 // first color in palette
		blackIndex = 1 // next color in palette
	)
	anim := gif.GIF{LoopCount: int(nframes)}
	phase := 1.0 // phase difference
	for i := 0; i < int(nframes); i++ {
		rect := image.Rect(0, 0, int(2*size+1), int(2*size+1))
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t*freq + phase)
			y := math.Sin(t + phase)
			img.SetColorIndex(int(size+x*size), int(size+y*size), blackIndex)
		}
		phase += 1
		anim.Delay = append(anim.Delay, int(delay))
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

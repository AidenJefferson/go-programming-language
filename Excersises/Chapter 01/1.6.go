// This program creates a multi-coloured gif (run using >output.gif argument)
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0, 0, 128, 1}, color.RGBA{255, 255, 0, 1}, color.White, color.RGBA{0, 128, 0, 1}}

const (
	blackIndex  = 0 // first color in palette
	blueIndex   = 1 // next color in palette
	yellowIndex = 2 // next color in palette
	whiteIndex  = 3 // next color in palette
	greenIndex  = 4 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 8      // number of complete x oscillator revolutions
		res     = 0.0001 // angular resolution
		size    = 150    // image canvas covers [-size..+size]
		nframes = 68     // number of animation frames
		delay   = 8      // delay between frames in 10ms units
	)
	freq := rand.Float64() * 4.5 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 1.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t*freq + phase)
			y := math.Sin(t + phase)
			color := rand.Intn(4)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(color))
		}
		phase += 1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

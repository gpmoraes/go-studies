package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

// lissajous generates a GIF animation of a Lissajous figure and writes it to the given writer.
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Number of complete x oscillator revolutions
		res     = 0.001 // Angular resolution
		size    = 100   // Image canvas covers [-size..+size]
		nframes = 64    // Number of animation frames
		delay   = 8     // Delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // Relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // Rectangle to hold the image
		img := image.NewPaletted(rect, palette)      // Create a new image with the specified palette
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex) // Set pixel color in the image
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay) // Set the delay between frames
		anim.Image = append(anim.Image, img)   // Add the image to the animation frames
	}
	gif.EncodeAll(out, &anim) // Encode the animation frames and write to the output writer
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // Initialize the random number generator with the current time
	lissajous(os.Stdout)                   // Generate and output the Lissajous animation to the standard output
}

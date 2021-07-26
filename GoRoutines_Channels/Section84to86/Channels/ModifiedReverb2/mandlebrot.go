package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"sync"
	"time"
)

const N = 128

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	start := time.Now()
	// partition jobs
	size := height / N
	var wg sync.WaitGroup
	var done = make(chan struct{})
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for py := size * i; py < size*(i+1); py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			done <- struct{}{}
		}(i)
		//png.Encode(os.Stdouz, img) // NOTE: ignoring errors
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	for range done {
	}

	fmt.Printf("elapsed %fms\n", time.Since(start).Seconds()*1000)
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

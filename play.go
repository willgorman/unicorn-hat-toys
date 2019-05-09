package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	perlin "github.com/aquilax/go-perlin"
	unicorn "github.com/arussellsaw/unicorn-go"
	util "github.com/arussellsaw/unicorn-go/util"
)

// display scrolling perlin noise map
func main() {
	noise := perlin.NewPerlinRandSource(2, 2, 3, rand.NewSource(time.Now().UnixNano()))

	c := unicorn.Client{Path: unicorn.SocketPath}
	err := c.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	currentX := 0
	ticker := time.NewTicker(500 * time.Millisecond)

	c.SetBrightness(10)
	go func() {
		for range ticker.C {
			m := perlinMatrix(2, currentX, *noise)
			c.SetAllPixels(unicorn.DeMatrix(m))
			c.Show()
			currentX++
		}
	}()

	time.Sleep(2 * time.Minute)
	ticker.Stop()
	c.Clear()
	c.Show()

	// for x := 0.; x < 8; x++ {
	// 	for y := 0.; y < 8; y++ {
	// 		n := noise.Noise2D(x/10, y/10)
	// 		c.SetPixel(uint(x), uint(y), uint(mapToColor(n)), 0, uint(mapToColor(n)))
	// 	}
	// }

}

func perlinMatrix(originX, originY int, p perlin.Perlin) util.Matrix {
	m := util.Matrix{}

	for x := 0.; x < 8; x++ {
		perlinX := x + float64(originX)
		for y := 0.; y < 8; y++ {
			perlinY := y + float64(originY)
			noise := p.Noise2D(normalize(perlinX), normalize(perlinY))
			noise = math.Abs(noise + 0.5)
			if noise < 0.33 {
				m[uint(x)][uint(y)] = util.Blue
			} else if noise < 0.66 {
				m[uint(x)][uint(y)] = util.Green
			} else {
				m[uint(x)][uint(y)] = util.Red
			}
		}
	}

	return m
}

func normalize(v float64) float64 {
	normalizer := math.Ceil(math.Log10(v)) * 10
	return v / normalizer
}

func mapToColor(noise float64) float64 {
	return math.Abs(noise) * 255
}

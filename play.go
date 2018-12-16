package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	perlin "github.com/aquilax/go-perlin"
	unicorn "github.com/arussellsaw/unicorn-go"
	"github.com/arussellsaw/unicorn-go/util"
)

func main() {
	noise := perlin.NewPerlinRandSource(2, 2, 3, rand.NewSource(time.Now().UnixNano()))

	c := unicorn.Client{Path: unicorn.SocketPath}
	err := c.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	c.SetBrightness(30)
	for x := 0.; x < 8; x++ {
		for y := 0.; y < 8; y++ {
			n := noise.Noise2D(x/10, y/10)
			c.SetPixel(uint(x), uint(y), uint(mapToColor(n)), 0, uint(mapToColor(n)))
		}
	}

	c.Show()
}

func perlinMatrix(originX, originY int, p perlin.Perlin) util.Matrix {
	m := util.Matrix{}

	for x := 0.; x < 8; x++ {
		perlinX := x + float64(originX)
		for y := 0.; y < 8; y++ {
			perlinY := y + float64(originY)
			m[uint(x)][uint(y)].B = uint(mapToColor(p.Noise2D(normalize(perlinX), normalize(perlinY))))
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

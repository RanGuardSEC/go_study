package main

import (
	"fmt"
	"math"
)

type VER struct {
	X, Y float64
}

func Abs(u VER) float64 {
	return math.Sqrt(u.X*u.X + u.Y*u.Y)
}

func main() {
	v := VER{3, 4}
	fmt.Println(Abs(v))
}

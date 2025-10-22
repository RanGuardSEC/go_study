package main

import (
	"fmt"
	"math"
)

type Verte struct {
	X, Y float64
}

func Ab(v Verte) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Verte, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Verte{3, 4}
	Scale(&v, 10)
	fmt.Println(Ab(v))
}

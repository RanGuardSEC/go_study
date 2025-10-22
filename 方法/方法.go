package main

import (
	"fmt"
	"math"
)

type vvv struct {
	X, Y float64
}

func (v vvv) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := vvv{3, 4}
	fmt.Println(v.Abs())
}

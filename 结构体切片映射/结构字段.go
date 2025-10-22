package main

import (
	"fmt"
)

type Vert struct {
	X, Y int
}

func main() {
	v := Vert{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

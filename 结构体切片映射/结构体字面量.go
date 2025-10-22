package main

import "fmt"

type Versex struct {
	X, Y int
}

var (
	v1 = Versex{1, 2}
	v2 = Versex{X: 1}
	v3 = Versex{}
	p  = &Versex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

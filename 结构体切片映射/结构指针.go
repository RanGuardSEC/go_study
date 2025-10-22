package main

import "fmt"

type Ve struct {
	X, Y int
}

func main() {
	v := Ve{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

}

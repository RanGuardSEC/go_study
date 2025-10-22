package main

import "fmt"

type Vx struct {
	Lat, Long float64
}

// 若顶层类型只是一个类型名，那么你可以在字面量的元素中省略它
var mmm = map[string]Vx{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(mmm)
}

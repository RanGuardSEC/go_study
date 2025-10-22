package main

import "fmt"

type Verte struct {
	Lat, Long float64
}

var m map[string]Verte

func main() {
	m = make(map[string]Verte)
	m["Bell Labs"] = Verte{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

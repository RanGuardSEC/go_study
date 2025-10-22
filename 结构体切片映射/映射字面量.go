package main

import "fmt"

type x struct {
	Lat, Long float64
}

// 映射的字面量和结构体类似，只不过必须有键名。
var t = map[string]x{
	"Bell Labs": x{
		40.68433, -74.39967,
	},
	"Google": x{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(t)
}

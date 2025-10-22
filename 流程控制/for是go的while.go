package main

import "fmt"

func main() {
	sum := 1
	p := 1
	for sum < 1000 {
		sum += sum
		p = p + 1
	}
	fmt.Println(sum)
	fmt.Println(p)
}

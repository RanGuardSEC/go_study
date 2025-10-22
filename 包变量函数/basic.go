package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("=== Zero Values in Goload ===")
	fmt.Printf("int:%d\n", i)
	fmt.Printf("float64:%f\n", f)
	fmt.Printf("bool:%t\n", b)
	fmt.Printf("string:%q\n", s)
	i = 42
	f = 3.1415
	b = true
	s = "Goload"
	fmt.Println("\n=== After Initialization ===")
	fmt.Printf("int:%d\n", i)
	fmt.Printf("float64:%.4f\n", f)
	fmt.Printf("bool:%t\n", b)
	fmt.Printf("string:%q\n", s)
}

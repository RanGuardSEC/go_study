package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	//a[low : high],它会选出一个半闭半开区间，包括第一个元素，但排除最后一个元素
	fmt.Println(s)
}

package main

import "fmt"

func main() {
	var s []int
	printSl(s)
	//append 的第一个参数 s 是一个元素类型为 T 的切片，其余类型为 T 的值将会追加到该切片的末尾。
	// 可在空切片上追加
	s = append(s, 0)
	printSl(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSl(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4, 5, 6, 7, 8)
	printSl(s)
}

func printSl(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printS("a", a)
	//切片可以用内置函数 make 来创建，这也是创建动态数组的方式
	//make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
	b := make([]int, 0, 5)
	//第一个数指定长度（0），第二个数指定容量（5）
	printS("b", b)

	c := b[:2]
	printS("c", c)

	d := c[2:5]
	printS("d", d)
}

func printS(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

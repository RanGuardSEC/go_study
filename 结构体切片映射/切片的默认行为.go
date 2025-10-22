package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	//切片下界的默认值为 0，上界则是该切片的长度。
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)

}

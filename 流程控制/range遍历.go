package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// 遍历：数据处理算法等作用
func main() {
	var a int
	a = 2
	//for 循环的 range 形式可遍历切片或映射
	//当使用 for 循环遍历切片时，每次迭代都会返回两个值。
	//第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
	for i, v := range pow {
		v = v * a
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

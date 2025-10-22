package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 创建二维切片
	picture := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		// 分配每一行
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// 尝试不同的函数来生成不同的图像
			row[x] = uint8(x * y) // 修改这里尝试不同函数
			// row[x] = uint8((x + y) / 2)
			// row[x] = uint8(x ^ y)
			// row[x] = uint8(x % (y + 1))
		}
		picture[y] = row
	}

	return picture
}

func main() {
	pic.Show(Pic)
}

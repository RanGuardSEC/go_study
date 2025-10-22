package main

import "fmt"

func main() {
	m := make(map[string]int)
	//在映射 m 中插入或修改元素
	//m[key] = elem
	m["答案"] = 42
	fmt.Println("值：", m["答案"])
	//获取元素：
	//elem = m[key]
	m["答案"] = 48
	fmt.Println("值：", m["答案"])
	//删除元素：
	//delete(m, key)
	delete(m, "答案")
	fmt.Println("值：", m["答案"])
	//通过双赋值检测某个键是否存在：
	//elem, ok = m[key]
	//若 key 在 m 中，ok 为 true ；否则，ok 为 false。
	//若 key 不在映射中，则 elem 是该映射元素类型的零值。
	v, ok := m["答案"]
	fmt.Println("值：", v, "是否存在？", ok)
	//若 elem 或 ok 还未声明，可以使用短变量声明：
	//elem, ok := m[key]
}

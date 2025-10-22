package main

import "fmt"

type Vert struct {
	X, Y float64
}

func (v *Vert) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vert, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vert{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vert{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

//比较前两个程序，会注意到带指针参数的函数必须接受一个指针：
//var v Vertex
//ScaleFunc(v, 5)  // 编译错误！
//ScaleFunc(&v, 5) // OK
//而接收者为指针的的方法被调用时，接收者既能是值又能是指针：
//var v Vertex
//v.Scale(5)  // OK
//p := &v
//p.Scale(10) // OK
//对于语句 v.Scale(5) 来说，即便 v 是一个值而非指针，带指针接收者的方法也能被直接调用。
//也就是说，由于 Scale 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。

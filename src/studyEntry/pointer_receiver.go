package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//func (v Vertex) Scale(f float64) {	//不使用指针形态，是值传递，退出函数后v不会改变
func (v *Vertex) Scale(f float64) {		//Methods with pointer receivers can modify the value to which the receiver points
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func ptr_rec() {


	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)	// methods with pointer receivers take either a value or a pointer as the receiver
	ScaleFunc(p, 8)

	fmt.Println(v, p)

}

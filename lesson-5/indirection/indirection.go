package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)

	fmt.Println("---------------------")

	v2 := Vertex{3, 4}
	fmt.Println(v2.Abs())
	fmt.Println(AbsFunc(v2))

	p2 := &Vertex{4, 3}
	fmt.Println(p2.Abs())
	fmt.Println(AbsFunc(*p2))
}

//NOTLAR:
//nesne büyükse pointer kullan struct yani nesne dediğim minimal
//metodlardan en az biri veriyi değiştiriyorsa (mutasyon) hepsini pointer alacak şekilde yap farklılık olmasın

//GENEL KULLANIM;
//bir tip üzerindeki tüm metodların ya value ya da pointer receiver olması gerekir, çorba yapılmamalı.
//**GENELDE pointer receiver kullanılıyor**

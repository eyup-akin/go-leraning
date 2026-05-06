package main

import (
	"fmt"
	"math"
)

//

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println("---------------------")

	v2 := Vertex{3, 4}
	fmt.Println(Abs(v2))

	fmt.Println("---------------------")

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

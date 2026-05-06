package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	q  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})

	fmt.Println("--------------")

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	fmt.Println("--------------")

	v5 := Vertex{1, 2}
	p := &v5
	p.X = 1e9
	fmt.Println(v5)

	fmt.Println("--------------")

	fmt.Println(v1, q, v2, v3)

}

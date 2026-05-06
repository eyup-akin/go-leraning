package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

//burada bu metod direkt bizim interfaceden oluyor
//M() kullandığı için ekstradan birşey yazamaya gerek kalmıyor.

/*

func (t T) M() {
	fmt.Println(t.S)
}

*/

func main() {

	/*
		var i1 I = T{"hello"}
		i1.M()
	*/

	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

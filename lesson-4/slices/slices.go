package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var v []int = primes[1:4]
	fmt.Println(v)

	fmt.Println("---------------------")

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)

	y[0] = "XXX"
	fmt.Println(x, y)
	fmt.Println(names)

	fmt.Println("---------------------")

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)

	fmt.Println("---------------------")

	sv := []int{2, 3, 5, 7, 11, 13}

	sv = sv[1:4]
	fmt.Println(sv)

	sv = sv[:2]
	fmt.Println(sv)

	sv = sv[1:]
	fmt.Println(sv)

	fmt.Println("---------------------")

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice1(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice1(s)

	// Extend its length.
	s = s[:4]
	printSlice1(s)

	// Drop its first two values.
	s = s[2:]
	printSlice1(s)

	fmt.Println("---------------------")

	var t []int

	fmt.Println(t, len(t), cap(t))

	if t == nil {
		fmt.Println("nil!")
	}

	fmt.Println("---------------------")

	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)

	printSlice1(d)

}

func printSlice1(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

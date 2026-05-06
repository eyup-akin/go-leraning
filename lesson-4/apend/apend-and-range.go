package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	fmt.Println("----------------")

	for index, value := range pow { //burada range yi akümülatör gibi düşün, ilk elemanında indexi atıyor ikinci elemanında valueyi atıyor
		fmt.Printf("2**%d = %d\n", index, value)
	}

	fmt.Println("----------------")

	powi := make([]int, 10)
	for i := range powi {
		powi[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range powi {
		fmt.Printf("%d\n", value)
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

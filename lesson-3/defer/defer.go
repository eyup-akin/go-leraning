package main

//defer = ertele

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("------------------")

	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")

}

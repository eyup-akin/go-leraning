package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	var k int = 0
	sum2 := 1
	for sum2 < 1000 { //for ; sum2 < 1000; {}
		sum2 += sum2
		k++
	}
	fmt.Println(sum2, k)

	//while gibi kullanımı
	tom := 1
	for tom < 1000 {
		tom += tom
	}

	fmt.Println(tom)
}

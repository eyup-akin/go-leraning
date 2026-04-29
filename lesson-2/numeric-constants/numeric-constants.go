package main

import "fmt"

const (
	Big = 1 << 100
	//ifadesi, 1 rakamının yanına 100 tane sıfır koymak demek

	Small = Big >> 99
	//başlangıçta kesin bir tipi yoktur.içinde bulundukları bağlama göre şekil alır
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

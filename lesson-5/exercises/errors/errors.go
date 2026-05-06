package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//e yi önce floata cast ederek sonsuz döngü olayını çözüyoruz
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// negatifse error
		return 0, ErrNegativeSqrt(x)
	}
	//değilse cevap
	return math.Sqrt(x), nil

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

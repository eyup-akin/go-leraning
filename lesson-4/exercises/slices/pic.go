package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	//önce dy uzunluğunda satır sayısı oluşturyoruz
	picture := make([][]uint8, dy)

	for y := range picture {
		//her satır için dx uzunluğunda iç dilim oluşturuyoz sutun sayısı

		picture[y] = make([]uint8, dx)

		for x := range picture[y] {
			//her hücreyi formülle dolduruyoruz
			picture[y][x] = uint8(x ^ y)
		}
	}

	return picture
}

func main() {
	pic.Show(Pic)
}

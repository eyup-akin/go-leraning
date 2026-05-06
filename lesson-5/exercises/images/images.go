package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// kkendi Image tipimizi tanımlıyoudz
type Image struct {
	width, height int
}

// colorModel metodunu implement ediyoruz
// standart RGBA modelini kullanacağımızı belirtiyoruz
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// bounds metodunu implement ediyoruz
// resmin sınırını (0,0) dan (width, heigght) a kadar tanımlıyoroz
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

// at metodunu implement ediyoruz
// her pixel için bir formul gösterir
func (i Image) At(x, y int) color.Color {
	// v değeri piksellerin parlaklığını belirleyecek (0-255 arası)
	v := uint8(x ^ y) // Daha önce kullandığın x*y veya x^y formülü
	return color.RGBA{v, v, 255, 255}
}

func main() {
	// 256x256 boyutlarında bir resim objesi oluşturup sisteme veriyoruz.
	m := Image{256, 256}
	pic.ShowImage(m)
}

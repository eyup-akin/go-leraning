package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	//önce içteki asıl okuyucudan ham veriyi çekiyoruz
	n, err = rot.r.Read(p)

	//n kadar veri okundu, şimdi bu p slicesi içindeki verileri dnüştür
	for i := 0; i < n; i++ {
		b := p[i]

		//büyük harf kontrolu
		if b >= 'A' && b <= 'Z' {
			p[i] = (b-'A'+13)%26 + 'A'
		} else if b >= 'a' && b <= 'z' {
			//kucuk harf kontrou
			p[i] = (b-'a'+13)%26 + 'a'
		}
	}

	//okunan byte sayısını ve hatayı döndürüyoruz
	return n, err

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // "You cracked the code!"
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

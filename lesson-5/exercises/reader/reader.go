package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// go nun stantart readerini kullanmak için bu şekilde yazılmalı

func (r MyReader) Read(b []byte) (int, error) {
	//gelen slicenin (b) uzunluğu ne kadar sa o kadar "A" koyuyoruz
	for i := range b {
		b[i] = 'A'
	}

	//slicein uzunluğu kadar veri yazdık ve hata ok döndür
	return len(b), nil

}

func main() {
	//reader.Validate fonk yazdığım readerin
	//sonsuza kadar 'A' üretip üretmediğini kontrol eder
	reader.Validate(MyReader{})
}

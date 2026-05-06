package main

import (
	"strings" // fonk için ekledik

	"golang.org/x/tour/wc" //uzak sever için
)

func WordCount(s string) map[string]int {
	//kelimeleri tutacak boş bir map oluşturuyoruz
	counts := make(map[string]int)

	//cümleyi kelimelere parçalıyoruz slice döner
	words := strings.Fields(s)

	//kelimelerin üzerinde dönüyoruz
	for _, word := range words {
		//her kelimeyi gördüğümüzde 1 arttırıyoruz countu
		//go default olarak 0 dan başlatıyo
		counts[word]++
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}

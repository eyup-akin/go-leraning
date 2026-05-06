package main

import (
	"fmt"
	"strings"
)

func main() {
	metin := "go çok hızlı ama go öğrenmek emek ister emek olmadan go olmaz"

	//kelimeleri ayırmak için map oluşturuyoruz
	//kelime string kav kere gectiği int

	sayac := make(map[string]int)

	//metni boşluklardan ayırıp bir slice yapıyoruz
	kelimeler := strings.Fields(metin)

	//her kelime üzerinde dönüyoruz
	for _, kelime := range kelimeler {
		//burada aslında kelime daha önceden eklenmemişse fekansına
		//ilk önce 0 yapmamız gerekebilir ama go otomatik yapıyor bunu

		sayac[kelime]++
	}

	fmt.Println("--- Kelime İstatistikleri ---")
	for kelime, adet := range sayac {
		fmt.Printf("'%s' kelimesi %d kez geçiyor.\n", kelime, adet)
	}

	//özel olarak bir kelimeyi kontrol edelim
	kontrol := "go"

	if adet, ok := sayac[kontrol]; ok {
		fmt.Printf("\nListedeki %s sayısı: %d", kontrol, adet)
	} else {
		fmt.Printf("\n'%s' kelimesi metinde hiç geçmiyor!", kontrol)
	}
}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// değer kopyalayarak
func yaslandirKopya(p Person) {
	p.Age += 10
	fmt.Println("Fonksiyon içi (Kopya):", p.Age)
}

// pointer ile alan gerçek değer değiştiren
func yaslandirGercek(p *Person) {
	//go nun minimal güzelliği (*p).Age yazmadan p.Age şeklinde
	//erişebiliyorsun pointer üzerinden struct'a
	p.Age += 10
	fmt.Println("Fonksiyon içi (Gerçek):", p.Age)
}

func main() {
	p1 := Person{Name: "Eyüp", Age: 20}

	fmt.Println("Başlangıç Yaşı::", p1.Age)

	//senaryo 1: kopya göndererek minimal
	yaslandirKopya(p1)
	fmt.Println("yaslandirKopya sonrası ana program: ", p1.Age)

	fmt.Println("-------------------")

	//senaryo 2: adres & gödererek minimal

	yaslandirGercek(&p1)

	fmt.Println("yaslandirGercek sonrası ana program: ", p1.Age)
}

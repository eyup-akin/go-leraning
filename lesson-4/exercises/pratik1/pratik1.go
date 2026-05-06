package main

import "fmt"

// struck tanımlama
type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {

	//person tipinde bir slice oluştur
	people := []Person{
		{Name: "Eyüp", Age: 20, Email: "bhdfdhf.com"},
		{Name: "Junior", Age: 1, Email: "dlhdhfd.com"},
	}

	//apend ile yeni kişiler ekleme

	people = append(people, Person{"Ahmet", 25, "ahmet@gmail.com"})
	people = append(people, Person{"Zeynep", 22, "zeynep@hotmail.com"})
	people = append(people, Person{"Can", 30, "can@pau.edu.tr"})

	//garip slice söngüsü
	fmt.Println("----- Kayıtlı Kişiler -----")
	for i, person := range people {
		fmt.Printf("%d. İsim: %s, Yaş: %d, E-posta: %s\n", i+1, person.Name, person.Age, person.Email)
	}

}

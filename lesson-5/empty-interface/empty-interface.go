package main

import "fmt"

func main() {
	var u interface{}
	describe(u)

	u = 42
	describe(u)

	u = "hello"
	describe(u)

	fmt.Println("***************")

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/*

		f = i.(float64) // panic modu aktif olur biladerim.
		fmt.Println(f)

	*/

	fmt.Println("***************")

	do(21)
	do("hello")
	do(true)
}

// buradaki "type" anahtar kelimesi sadece switch ile kullanabilinen özel bir yapı
// go burada type bölmesini okur ve ona denk gelen case'i yapar biladerim.
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

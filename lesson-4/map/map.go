package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var maps = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var maps2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {

	fmt.Println("--------------------")

	fmt.Println(maps)

	fmt.Println("--------------------")

	fmt.Println(maps2)

	fmt.Println("--------------------")

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] //ilki değer ikinci var-yok
	fmt.Println("The value:", v, "Present?", ok)
}

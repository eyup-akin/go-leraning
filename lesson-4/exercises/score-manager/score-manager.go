package main

import "fmt"

func scoreManager(playerName string) func(points int) {
	score := 0 //bu bizim gizli hafıza

	return func(points int) {
		score += points
		fmt.Printf("%s adlı oyuncunun güncel skoru: %d\n", playerName, score)
	}
}

func main() {
	//iki ayrı hafıza alanı oluşturuyooruz.
	player1 := scoreManager("Eyüp")
	player2 := scoreManager("Junior-Bot")

	player1(10)
	player1(5)
	player2(20)
}

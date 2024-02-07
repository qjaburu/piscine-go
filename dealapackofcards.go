package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	rangDeck := deck // result of deck
	Player := 1
	for i := '0'; i < 12; i += 3 {

		fmt.Printf("Player %v: %v,%v, %v\n", Player, rangDeck[i], rangDeck[i+1], rangDeck[i+2])
		Player++
	}
}

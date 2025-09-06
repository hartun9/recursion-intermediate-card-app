package main

import (
	"fmt"

	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
)

func main() {
	dealer := cards.Dealer{}
	game1 := dealer.StartGame(4)
	for _, card := range game1[0] {
		fmt.Println(card.GetCardString())
	}
}

package main

import (
	"fmt"

	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
)

func main() {
	dealer := cards.Dealer{}
	table1 := cards.Table{AmountOfPlayers: 4}
	game1 := dealer.StartGame(table1)
	for _, card := range game1[0] {
		fmt.Println(card.GetCardString())
	}
}

package main

import (
	"fmt"

	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
)

func main() {
	dealer := cards.Dealer{}
	table1 := dealer.StartGame(4)
	for _, card := range table1[0] {
		fmt.Println(card.GetCardString())
	}
}

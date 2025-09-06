package main

import (
	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
)

func main() {
	dealer := cards.Dealer{}

	table2 := cards.Table{AmountOfPlayers: 4, GameMode: "poker"}
	game2 := dealer.StartGame(table2)
	dealer.PrintTableInformation(game2, table2)
}

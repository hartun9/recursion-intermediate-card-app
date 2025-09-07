package main

import (
	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
	"github.com/hartun9/recursion-intermediate-card-app/internal/dealer"
)

func main() {
	dealer := dealer.Dealer{}

	table1 := cards.Table{AmountOfPlayers: 3, GameMode: "21"}
	game1 := dealer.StartGame(table1)
	dealer.PrintTableInformation(game1, table1)

	table2 := cards.Table{AmountOfPlayers: 4, GameMode: "poker"}
	game2 := dealer.StartGame(table2)
	dealer.PrintTableInformation(game2, table2)
}

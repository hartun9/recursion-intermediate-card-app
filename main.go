package main

import (
	"fmt"

	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
	"github.com/hartun9/recursion-intermediate-card-app/internal/dealer"
)

func main() {
	table1 := cards.Table{AmountOfPlayers: 4, GameMode: "21"}
	game1 := dealer.StartGame(table1)
	dealer.PrintTableInformation(game1, table1)

	fmt.Println(dealer.WinnerOf21(game1))
}

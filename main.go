package main

import (
	"fmt"

	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
	"github.com/hartun9/recursion-intermediate-card-app/internal/dealer"
)

func main() {
	table1 := cards.NewTable(1, "poker")
	game1 := dealer.StartGame(*table1)

	table2 := cards.NewTable(3, "21")
	game2 := dealer.StartGame(*table2)

	table3 := cards.NewTable(2, "pair_of_cards")
	game3 := dealer.StartGame(*table3)

	dealer.PrintTableInformation(game1, *table1)
	fmt.Println(dealer.CheckWinner(game1, *table1))

	dealer.PrintTableInformation(game2, *table2)
	fmt.Println(dealer.CheckWinner(game2, *table2))

	dealer.PrintTableInformation(game3, *table3)
	fmt.Println(dealer.CheckWinner(game3, *table3))
}

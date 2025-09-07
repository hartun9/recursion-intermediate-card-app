package main

import (
	"fmt"

	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
	"github.com/hartun9/recursion-intermediate-card-app/internal/dealer"
)

func main() {
	playerA := []*cards.Card{
		{"♦︎", "A", 1},
		{"♦︎", "J", 11},
	}
	playerB := []*cards.Card{
		{"♦︎", "9", 9},
		{"♦︎", "K", 13},
	}

	fmt.Println(dealer.Score21Individual(playerA))
	fmt.Println(dealer.Score21Individual(playerB))
}

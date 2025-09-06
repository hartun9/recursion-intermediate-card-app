package cards

import (
	"fmt"
	"strconv"
)

type Dealer struct{}

// StartGame 参加人数を受け取り、それぞれのプレイヤーにカードを配る
func (Dealer) StartGame(table Table) [][]*Card {
	deck := NewDeck()
	deck.ShuffleDeck()
	var playerCards [][]*Card
	for i := 0; i < table.AmountOfPlayers; i++ {
		var playerHand []*Card
		dealer := Dealer{}
		for j := 0; j < dealer.InitialCards(table.GameMode); j++ {
			card1 := deck.Draw()
			playerHand = append(playerHand, card1)
		}
		playerCards = append(playerCards, playerHand)
	}
	return playerCards
}

// InitialCards gameModeに応じて手札の数を返す
func (Dealer) InitialCards(gameMode string) int {
	switch gameMode {
	case "21":
		return 2
	case "poker":
		return 5
	}
	return 0
}

func (Dealer) PrintTableInformation(playerCards [][]*Card, table Table) {
	fmt.Println("Amount of players: " + strconv.Itoa(table.AmountOfPlayers) + "... Game mode: " + table.GameMode + ". At this table: ")

	for i, playerCard := range playerCards {
		fmt.Println("Player " + strconv.Itoa(i+1) + " hand is: ")
		for _, card := range playerCard {
			fmt.Print(card.GetCardString())
		}
		fmt.Println()
	}
}

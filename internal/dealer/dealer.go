package dealer

import (
	"fmt"
	"strconv"

	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
	"github.com/hartun9/recursion-intermediate-card-app/internal/game"
	"github.com/hartun9/recursion-intermediate-card-app/internal/helper"
)

// StartGame 参加人数を受け取り、それぞれのプレイヤーにカードを配る
func StartGame(table cards.Table) [][]*cards.Card {
	deck := cards.NewDeck(table)
	deck.ShuffleDeck()
	var playerCards [][]*cards.Card
	for i := 0; i < table.AmountOfPlayers; i++ {
		var playerHand []*cards.Card
		for j := 0; j < InitialCards(table.GameMode); j++ {
			card1 := deck.Draw()
			playerHand = append(playerHand, card1)
		}
		playerCards = append(playerCards, playerHand)
	}
	return playerCards
}

// InitialCards gameModeに応じて手札の数を返す
func InitialCards(gameMode string) int {
	switch gameMode {
	case "21":
		return 2
	case "poker":
		return 5
	case "pair_of_cards":
		return 5
	}
	return 0
}

func PrintTableInformation(playerCards [][]*cards.Card, table cards.Table) {
	fmt.Println("Amount of players: " + strconv.Itoa(table.AmountOfPlayers) + "... Game mode: " + table.GameMode + ". At this table: ")

	for i, playerCard := range playerCards {
		fmt.Println("Player " + strconv.Itoa(i+1) + " hand is: ")
		for _, card := range playerCard {
			fmt.Print(card.GetCardString())
		}
		fmt.Println()
	}
}

func Score21Individual(cards []*cards.Card) int {
	var value int
	for _, card := range cards {
		value += card.IntValue
	}
	if value > 21 {
		return 0
	}
	return value
}

func WinnerOf21(playerCardsList [][]*cards.Card) string {
	points := make([]int, len(playerCardsList))
	pointCount := make(map[int]int)

	for i, playerCards := range playerCardsList {
		point := Score21Individual(playerCards)
		points[i] = point
		pointCount[point]++
	}

	winnerIndex := helper.MaxInArrayIndex(points)
	if pointCount[points[winnerIndex]] > 1 {
		return "It is a draw "
	} else if pointCount[points[winnerIndex]] >= 0 {
		return "player " + strconv.Itoa(winnerIndex+1) + " is the winner"
	} else {
		return "No winners.."
	}
}

func CheckWinner(playerCardsList [][]*cards.Card, table cards.Table) string {
	switch table.GameMode {
	case "21":
		return WinnerOf21(playerCardsList)
	case "pair_of_cards":
		return game.WinnerPairOfCards(playerCardsList[0], playerCardsList[1])
	default:
		return "no game"
	}
}

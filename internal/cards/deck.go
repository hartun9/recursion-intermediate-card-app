package cards

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	Deck []Card
}

func NewDeck(table Table) *Deck {
	return &Deck{Deck: *generateDeck(table)}
}

func numForGame(gameMode string, j int, value string) int {
	var gameNumMap map[string]int
	hasGameNumMap := false
	switch gameMode {
	case "21":
		gameNumMap = map[string]int{
			"A": 1,
			"J": 10,
			"Q": 10,
			"K": 10,
		}
		hasGameNumMap = true
	case "pair_of_cards":
		gameNumMap = map[string]int{
			"A": 14,
			"K": 13,
			"Q": 12,
			"J": 11,
		}
		hasGameNumMap = true
	}
	var intValue int
	if hasGameNumMap {
		if num, found := gameNumMap[value]; found {
			intValue = num
		} else {
			intValue = j + 1
		}
	} else {
		intValue = j + 1
	}
	return intValue
}

func generateDeck(table Table) *[]Card {
	var newDeck []Card
	suits := []string{"♣", "♦", "♥", "♠"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for j, value := range values {
			intValue := numForGame(table.GameMode, j, value)
			newDeck = append(newDeck, Card{Suit: suit, Value: value, IntValue: intValue})
		}
	}
	return &newDeck
}

func (d *Deck) PrintDeck() {
	fmt.Println("Displaying cards...")
	for _, c := range d.Deck {
		fmt.Println(c.GetCardString())
	}
}

func (d *Deck) ShuffleDeck() {
	deckSize := len(d.Deck)
	for i := deckSize - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		temp := d.Deck[i]
		d.Deck[i] = d.Deck[j]
		d.Deck[j] = temp
	}
}

func (d *Deck) Draw() *Card {
	last := d.Deck[len(d.Deck)-1]
	d.Deck = d.Deck[:len(d.Deck)-1]
	return &last
}

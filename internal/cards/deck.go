package cards

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	Deck []Card
}

func NewDeck() *Deck {
	return &Deck{Deck: *generateDeck()}
}

func generateDeck() *[]Card {
	var newDeck []Card
	suits := []string{"♣", "♦", "♥", "♠"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for j, value := range values {
			newDeck = append(newDeck, Card{Suit: suit, Value: value, IntValue: j + 1})
			c := Card{Suit: suit, Value: value, IntValue: j + 1}
			fmt.Println(c.GetCardString())
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

package cards

type Dealer struct{}

// StartGame 参加人数を受け取り、それぞれのプレイヤーにカードを配る
func (Dealer) StartGame(table Table) [][]*Card {
	deck := NewDeck()
	deck.ShuffleDeck()
	var playerCards [][]*Card
	for i := 0; i < table.AmountOfPlayers; i++ {
		var playerHand []*Card
		for j := 0; j < 2; j++ {
			card1 := deck.Draw()
			playerHand = append(playerHand, card1)
		}
		playerCards = append(playerCards, playerHand)
	}
	return playerCards
}

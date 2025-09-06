package cards

type Dealer struct{}

// StartGame 参加人数を受け取り、それぞれのプレイヤーにカードを配る
func (Dealer) StartGame(amountOfPlayers int) [][2]*Card {
	deck := NewDeck()
	deck.ShuffleDeck()
	var playerCards [][2]*Card
	for i := 0; i < amountOfPlayers; i++ {
		var playerHand [2]*Card
		for j := 0; j < 2; j++ {
			card1 := deck.Draw()
			playerHand[j] = card1
		}
		playerCards = append(playerCards, playerHand)
	}
	return playerCards
}

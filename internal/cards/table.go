package cards

type Table struct {
	AmountOfPlayers int
	GameMode        string
}

func NewTable(amountOfPlayers int, gameMode string) *Table {
	if amountOfPlayers < 1 {
		panic("amount of players must be greater than zero")
	}
	switch gameMode {
	case "pair_of_cards":
		if amountOfPlayers != 2 {
			panic("cards table must contain two players")
		}
	}
	return &Table{AmountOfPlayers: amountOfPlayers, GameMode: gameMode}
}

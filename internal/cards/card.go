package cards

import "strconv"

type Card struct {
	Suit     string
	Value    string
	IntValue int
}

func (c *Card) GetCardString() string {
	return c.Suit + c.Value + "(" + strconv.Itoa(c.IntValue) + ")"
}

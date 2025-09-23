package game

import (
	"testing"

	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
)

func TestWinnerPairOfCards(t *testing.T) {
	tests := []struct {
		name string
		p1   []int
		p2   []int
		want string
	}{
		{
			name: "player 1 higher multiplicity wins",
			p1:   []int{4, 8, 8, 8, 11},
			p2:   []int{4, 11, 11, 12, 3},
			want: "player 1 is the winner",
		},
		{
			name: "ペアの枚数が同じで、player 2 の rank が高い",
			p1:   []int{4, 8, 8, 4, 11},
			p2:   []int{4, 11, 11, 12, 3},
			want: "player 2 is the winner",
		},
		{
			name: "ペアの枚数が同じで、player 1 の rank が高い",
			p1:   []int{4, 7, 7, 12, 11},
			p2:   []int{7, 7, 13, 12, 2},
			want: "player 2 is the winner",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WinnerPairOfCards(cardsFromInts(tt.p1...), cardsFromInts(tt.p2...))
			if got != tt.want {
				t.Fatalf("WinnerPairOfCards() = %q, want %q", got, tt.want)
			}
		})
	}
}

func cardsFromInts(values ...int) []*cards.Card {
	cardsSlice := make([]*cards.Card, 0, len(values))
	for _, v := range values {
		cardsSlice = append(cardsSlice, &cards.Card{IntValue: v})
	}
	return cardsSlice
}

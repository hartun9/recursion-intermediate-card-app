package game

import (
	"strconv"

	"github.com/hartun9/recursion-intermediate-card-app/internal/cards"
)

type PairOfCards struct {
}

type MaxRank struct {
	count int
	rank  int
}

type Player struct {
	rankCount map[int]int
	maxRank   MaxRank
}

func WinnerPairOfCards(player1 []*cards.Card, player2 []*cards.Card) string {
	var cards1 []string
	var cards2 []string

	for _, card := range player1 {
		cards1 = append(cards1, card.Suit+card.Value)
	}

	for _, card := range player2 {
		cards2 = append(cards2, card.Suit+card.Value)
	}

	p1 := newPlayer(cards1)
	p2 := newPlayer(cards2)

	return winnerPairOfCardsHelper(p1, p2)
}

func newPlayer(playerCards []string) *Player {
	rankCount := newRankCount(playerCards)
	return &Player{
		rankCount: rankCount,
		maxRank:   calculateMaxRankCount(rankCount),
	}
}

func nextPlayer(p *Player) *Player {
	delete(p.rankCount, p.maxRank.rank)
	return &Player{
		rankCount: p.rankCount,
		maxRank:   calculateMaxRankCount(p.rankCount),
	}
}

func winnerPairOfCardsHelper(p1 *Player, p2 *Player) string {
	if len(p1.rankCount) == 0 || len(p2.rankCount) == 0 {
		return "It is a draw "
	}

	if p1.maxRank.count > p2.maxRank.count {
		return "player 1 is the winner"
	} else if p2.maxRank.count > p1.maxRank.count {
		return "player 2  is the winner"
	} else {
		if p1.maxRank.rank > p2.maxRank.rank {
			return "player 1 is the winner"
		} else if p1.maxRank.rank == p2.maxRank.rank {
			return winnerPairOfCardsHelper(nextPlayer(p1), nextPlayer(p2))
		}
		return "player 2 is the winner"
	}
}

func newRankCount(playerCards []string) (rankCount map[int]int) {
	rankCount = make(map[int]int)
	for _, card := range playerCards {
		cardRune := []rune(card)
		rank := string(cardRune[1:])
		rankNum := 0
		switch rank {
		case "A":
			rankNum = 14
		case "K":
			rankNum = 13
		case "Q":
			rankNum = 12
		case "J":
			rankNum = 11
		default:
			rankNum, _ = strconv.Atoi(rank)
		}
		rankCount[rankNum]++
	}
	return rankCount
}

func calculateMaxRankCount(rankCount map[int]int) MaxRank {
	var maxRankCount int
	var maxRankCountRank int
	for rank, count := range rankCount {
		if count > maxRankCount {
			maxRankCount = count
			maxRankCountRank = rank
		} else if count == maxRankCount {
			if rank > maxRankCountRank {
				maxRankCountRank = rank
			}
		}
	}
	return MaxRank{maxRankCount, maxRankCountRank}
}

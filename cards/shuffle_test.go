package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	someCards := []Card{
		{
			Suit: "a",
			Face: "1",
		},
		{
			Suit: "b",
			Face: "2",
		},
		{
			Suit: "v",
			Face: "5",
		},
	}

	deck := FactoryDeck{
		Cards: someCards,
	}

	deck.Shuffle()

	assert.NotEqual(t, Card{"a", "1", false}, deck.Cards[0])
}

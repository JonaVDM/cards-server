package core

type Card struct {
	Suit  string
	Face  string
	Joker bool
}

type Deck interface {
	Shuffle()
}

type FactoryDeck struct {
	Cards []Card
}

func New(size int) Deck {
	suits := [4]string{
		"hearts",
		"spades",
		"diamond",
		"clubs",
	}

	faces := [13]string{
		"ace",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"jack",
		"queen",
		"king",
	}

	cards := make([]Card, 0)

	for i := 0; i < size; i++ {
		// Generate the deck
		for _, s := range suits {
			for _, f := range faces {
				cards = append(cards, Card{s, f, false})
			}
		}

		// Generate two jokers
		cards = append(cards, Card{"black", "", true})
		cards = append(cards, Card{"red", "", true})
	}

	return FactoryDeck{
		Cards: cards,
	}
}

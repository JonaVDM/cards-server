package cards

import "math/rand"

func (d FactoryDeck) Shuffle() {
	for x := 0; x < 100; x++ {
		for i := range d.Cards {
			j := rand.Intn(i + 1)
			d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
		}
	}
}

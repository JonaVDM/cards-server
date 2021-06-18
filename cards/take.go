package cards

func (d *FactoryDeck) Take() Card {
	card := d.Cards[(len(d.Cards) - 1)]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card
}

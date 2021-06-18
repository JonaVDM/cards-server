package cards

type Action struct {
	Rule string
	Data int
}

func (d *FactoryDeck) Play(card Card) Action {
	d.Discarded = append(d.Discarded, card)

	if card.Joker {
		return Action{
			Rule: "grab",
			Data: 5,
		}
	}

	switch card.Face {
	case "two":
		return Action{
			Rule: "grab",
			Data: 2,
		}

	case "seven":
		return Action{
			Rule: "again",
		}

	case "eight":
		return Action{
			Rule: "skip",
		}

	case "jack":
		return Action{
			Rule: "change",
		}
	}

	return Action{}
}

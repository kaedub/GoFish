package main

type BotMemory struct {
	// requests maps hold card number as key
	// and the turn numbers the requests took place as value
	requestsMade     map[int][]int
	requestsReceived map[int][]int
}

func (b *BotMemory) setRequestMade(cardNumber int, turn int) {
	b.requestsMade[cardNumber] = append(b.requestsMade[cardNumber], turn)
}
func (b *BotMemory) setRequestReceived(cardNumber int, turn int) {
	b.requestsReceived[cardNumber] = append(b.requestsReceived[cardNumber], turn)
}

func (b BotMemory) requestCard(hand Deck, turnCount int) int {
	selected := Card{}
	for _, card := range hand {
		if b.requestsMade[card.Value] != nil {
			distance := turnCount
			for _, turn := range b.requestsMade[card.Value] {
				if turnCount-turn < distance {
					distance = turnCount - turn
				}
			}
			if distance > 6 {
				selected = card
			}
		}
	}
	return selected.Value
}

package main

import (
	"testing"
)

func BotTest(t *testing.T) {
	bot := BotMemory{}

	turn := 20
	deck := newDeck()
	deck.shuffle()
	hand := Deck{}
	for i := 0; i < 5; i++ {
		card := deck.draw()
		hand = append(hand, card)
		if i > 0 {
			bot.setRequestMade(card.Value, 4)
			bot.setRequestMade(card.Value, 10)
		} else {
			bot.setRequestMade(card.Value, 8)
			bot.setRequestMade(card.Value, 16)
		}
	}

	requestedCardNumber := bot.requestCard(hand, turn)
	if requestedCardNumber != hand[0].Value {
		t.Errorf("bot.requestCard did not choose the correct card:\nexpected -> %v\nreceived -> %v", hand[0].Value, requestedCardNumber)
	}
}

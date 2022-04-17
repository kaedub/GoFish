package main

import (
	"testing"
)

func TestDeck(t *testing.T) {
	// assert deck contains only 52 cards
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expect %v cards but got %v cards", 52, len(deck))
	}

	// assert no duplicate cards are in deck
	counter := make(map[string]int)
	for _, card := range deck {
		if _, ok := counter[card.String()]; !ok {
			counter[card.String()] = 1
		} else {
			t.Errorf("Expect unique cards but found >1 of %v", card.String())
		}
	}

	// assert shuffle shuffles cards
	a, b, c, d := deck[0], deck[1], deck[2], deck[3]
	deck.shuffle()
	e, f, g, h := deck[0], deck[1], deck[2], deck[3]
	if a == e && b == f && c == g && d == h {
		t.Errorf("Shuffle did not shuffle cards")
	}

	// assert that draw removes a card from the top of the deck
	topCard := deck[len(deck)-1]
	drawnCard := deck.draw()
	if topCard != drawnCard {
		t.Errorf("deck.draw did not get the top card: %v vs %v", topCard, drawnCard)
	}
	if len(deck) != 51 {
		t.Errorf("deck.draw did not remove the top card")
	}

	// assert that remove can remove a card from the middle of the deck
	deck = newDeck()
	selectedCard := deck[5]
	removedCard := deck.remove(5)
	if selectedCard != removedCard {
		t.Errorf("deck.remove did not get the correct card: %v vs %v", selectedCard, removedCard)
	}
	if len(deck) != 51 {
		t.Errorf("deck.remove did not remove the card: %v vs %v", selectedCard, removedCard)
	}

	// assert that deck will sort cards by value
	deck = newDeck()
	deck.Sort()
	highCardValue := 14
	for i := 0; i < 12; i++ {
		if deck[i].Value != highCardValue-(i/4) {
			t.Errorf("deck.sort did not sort from highest to lowest")
		}
	}
}
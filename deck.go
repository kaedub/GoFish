package main

import (
	"math/rand"
	"sort"
	"time"
)

type Deck []Card

var cardValues = 14

func newDeck() Deck {
	var deck Deck

	for val := 2; val < cardValues+1; val++ {
		for _, suit := range suits {
			card := Card{
				Value: val,
				Suit:  suit,
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func (d Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d *Deck) draw() Card {
	index := len(*d) - 1
	card := (*d)[index]
	*d = (*d)[:index] // Remove it from the stack by slicing it off.
	return card
}

func (d *Deck) remove(index int) Card {
	card := (*d)[index]
	*d = append((*d)[:index], (*d)[index+1:]...)
	return card
}

func (d *Deck) removeMany(indices []int) Deck {
	indexSet := make(map[int]bool)
	for _, i := range indices {
		indexSet[i] = true
	}

	tempDeck := []Card{}
	removed := []Card{}
	for i, card := range *d {
		if indexSet[i] {
			removed = append(removed, card)
		} else {
			tempDeck = append(tempDeck, card)
		}
	}
	(*d) = tempDeck
	return removed
}

func (d *Deck) Sort() {
	sort.Slice((*d), func(i, j int) bool {
		return (*d)[i].Value > (*d)[j].Value
	})
}

func (d *Deck) removeMatches(minimumMatchCount int) Deck {
	m := make(map[string][]int)
	for i, card := range *d {
		if _, ok := m[card.ValueToString()]; !ok {
			m[card.ValueToString()] = []int{i}
		} else {
			m[card.ValueToString()] = append(m[card.ValueToString()], i)
		}
	}

	indicesToRemove := []int{}
	for _, v := range m {
		if len(v) >= minimumMatchCount {
			indicesToRemove = append(indicesToRemove, v...)
		}
	}

	return d.removeMany(indicesToRemove)
}

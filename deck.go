package main

import (
	"fmt"
	"math/rand"
	"sort"
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

func (d Deck) print() {
	for _, card := range d {
		fmt.Println(card.String())
	}
}

func (d Deck) shuffle() {
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(22)
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

func (d *Deck) Sort() {
	sort.Slice((*d), func(i, j int) bool {
		return (*d)[i].Value > (*d)[j].Value
	})
}
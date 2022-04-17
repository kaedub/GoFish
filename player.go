package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Player struct {
	IsBot bool
	Name  string
	Hand  Deck
}

func newPlayer(name string, isBot bool) Player {
	return Player{
		Name:  name,
		IsBot: isBot,
	}
}

func (p Player) showHand() string {
	var hand strings.Builder
	hand.WriteString(fmt.Sprintf("%v's Hand\n", p.Name))
	for _, card := range p.Hand {
		s := fmt.Sprintf("%v\n", card.String())
		hand.WriteString(s)
	}
	return hand.String()
}

func (p *Player) fish(reader *bufio.Reader) string {
	if p.IsBot {
		return p.Hand[0].ValueToString()
	} else {
		return GetInputFromUser("Enter a card value (2-10, J, Q, K, or A) to ask for", reader)
	}
}

func (p *Player) respond(s string) (Card, bool) {
	v := StringToValue(s)
	for i, card := range p.Hand {
		if card.Value == v {
			return p.Hand.remove(i), true
		}
	}
	return Card{}, false
}

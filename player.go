package main

import (
	"bufio"
)

type Player struct {
	Name  string
	Hand  Deck
	Books Deck
}

func newPlayer(name string) Player {
	return Player{
		Name: name,
	}
}

func (p *Player) addToHand(c Card) {
	p.Hand = append(p.Hand, c)
}

func (p *Player) setBooks() {
	books := p.Hand.removeMatches(4)
	books.Sort()
	p.Books = append(p.Books, books...)
}

func (p Player) getScore() int {
	score := 0
	for _, card := range p.Books {
		score += card.Value
	}
	return score
}

func (p *Player) userSeek(reader *bufio.Reader) string {
	return GetInputFromUser("Enter a card value (2-10, J, Q, K, or A) to ask for", reader)
}

func (p *Player) botSeek() string {
	// TODO: Implement bot logic
	return p.Hand[0].ValueToString()
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

func (p Player) endTurn(reader *bufio.Reader) {
	GetInputFromUser("[Enter] to end turn", reader)
}

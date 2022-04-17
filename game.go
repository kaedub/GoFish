package main

import (
	"bufio"
	"fmt"
)

type Game struct {
	Player   Player
	Opponent Player
	Deck     Deck
}

var getInput = GetInputFromUser

func newGame(reader *bufio.Reader) Game {
	name := GetInputFromUser("Enter your name:", reader)
	return Game{
		Player:   newPlayer(name, false),
		Opponent: newPlayer("CPU", true),
		Deck:     newDeck(),
	}
}

func (g Game) start(reader *bufio.Reader) {
	g.Deck.shuffle()
	g.deal()

	gameOver := false
	activePlayer := &g.Player
	inactivePlayer := &g.Opponent
	turn := 1
	for gameOver == false {
		fmt.Println()
		fmt.Printf("===============  TURN %v ==================\n", turn)
		fmt.Printf("It is %v's turn \n", activePlayer.Name)

		cardValue := activePlayer.fish(reader)
		fmt.Printf("%v: Do you have have any %vs?\n", activePlayer.Name, StringToFullString(cardValue))
		card, success := inactivePlayer.respond(cardValue)
		if success {
			fmt.Printf("%v gives up card %v\n", inactivePlayer.Name, card.String())
			(*activePlayer).Hand = append(activePlayer.Hand, card)
		} else {
			fmt.Printf("%v: Go Fish!\n", inactivePlayer.Name)
			(*activePlayer).Hand = append(activePlayer.Hand, g.Deck.draw())
		}

		g.Player.Hand.Sort()
		g.Opponent.Hand.Sort()

		fmt.Println()
		fmt.Println(g.Player.showHand())

		GetInputFromUser("[Enter] to end turn", reader)

		tmp := inactivePlayer
		inactivePlayer = activePlayer
		activePlayer = tmp
		turn++
	}
}

func (g *Game) deal() {
	for i := 0; i < 7; i++ {
		g.Player.Hand = append(g.Player.Hand, g.Deck.draw())
		g.Opponent.Hand = append(g.Opponent.Hand, g.Deck.draw())
	}
	g.Player.Hand.Sort()
	g.Opponent.Hand.Sort()
	fmt.Println(g.Player.showHand())
}

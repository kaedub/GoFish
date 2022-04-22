package main

import (
	"bufio"
	"fmt"
)

type Game struct {
	Player Player
	Bot    Player
	Deck   Deck
}

func newGame(reader *bufio.Reader) Game {
	name := GetInputFromUser("Enter your name:", reader)
	return Game{
		Player: newPlayer(name),
		Bot:    newPlayer("CPU"),
		Deck:   newDeck(),
	}
}

func (g Game) start(reader *bufio.Reader) {
	initHandSize := 7
	g.Deck.shuffle()
	g.deal(initHandSize)

	var gameOver = false
	turnCount := 1
	for !gameOver {
		// player turn
		g.turn(reader, false, turnCount)
		turnCount++

		// bot turn
		g.turn(reader, true, turnCount)
		turnCount++

		if turnCount > 20 {
			gameOver = true
		}
	}
}

func (g *Game) deal(handSize int) {
	cardsToDeal := handSize
	for ; cardsToDeal > 0; cardsToDeal-- {
		g.Player.addToHand(g.Deck.draw())
		g.Bot.addToHand(g.Deck.draw())
	}
	g.Player.Hand.Sort()
	g.Bot.Hand.Sort()
	fmt.Println(g.Player.Hand.String())
}

func (g *Game) turn(reader *bufio.Reader, bot bool, turnCount int) {
	requesting := &g.Player
	responding := &g.Bot
	if bot {
		requesting = &g.Bot
		responding = &g.Player
	}

	fmt.Println()
	fmt.Printf("===============  TURN %v ==================\n", turnCount)
	fmt.Printf("It is %v's turn \n", requesting.Name)

	var cardValue string
	if bot {
		cardValue = requesting.botSeek()
	} else {
		cardValue = requesting.userSeek(reader)
	}

	fmt.Printf("%v: Do you have have any %vs?\n", requesting.Name, StringToFullString(cardValue))
	card, success := responding.respond(cardValue)

	if success {
		fmt.Printf("%v gives up card %v\n", responding.Name, card.String())
		requesting.addToHand(card)
	} else {
		fmt.Printf("%v: Go Fish!\n", responding.Name)
		drawn := g.Deck.draw()
		if !bot {
			fmt.Printf("You drew a %v\n", drawn.String())
		}
		requesting.addToHand(drawn)
	}

	g.Player.setBooks()
	g.Bot.setBooks()
	fmt.Printf("Player books:\n%v\n", g.Player.Books.String())
	fmt.Printf("Player score: %v\n", g.Player.getScore())
	fmt.Printf("Bot books:\n%v\n", g.Bot.Books.String())
	fmt.Printf("Bot score %v\n", g.Bot.getScore())

	g.Player.Hand.Sort()
	g.Bot.Hand.Sort()

	fmt.Println()
	fmt.Println("Your Hand")
	fmt.Println(g.Player.Hand.String())

	g.Player.endTurn(reader)
}

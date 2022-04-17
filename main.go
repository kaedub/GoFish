package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	game := newGame(reader)
	game.start(reader)
}

package main

import "os"

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()

	os.Exit(0)
}
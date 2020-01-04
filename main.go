package main

import "fmt"

func main() {
	cards := newDeck()
	
	cards.shuffle()
	fmt.Println(deal(cards, 5))
	cards.print()
}

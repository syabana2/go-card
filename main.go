package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {

	cards := deck{newCard(), "One of Diamonds"}
	cards = append(cards, "Two of Heart")

	cards.print()

}

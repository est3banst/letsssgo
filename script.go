package main

func main() {
	cards := newDeck()

	hand, remainingCards := dealCards(cards, 3)

	// cards.print()
	hand.print()
	remainingCards.print()
}

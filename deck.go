package main

import "fmt"

type deck []string

// CREATE A DECK
func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{
		"Spades", "Diamonds", "Hearts", "Clubs",
	}
	cardsValues := []string{
		"Ace", "Two", "Three", "Four",
	}
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// PRINT CARDS
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// DEAL CARDS

func dealCards(d deck, hand int) (deck, deck) {
	return d[:hand], d[hand:]
}

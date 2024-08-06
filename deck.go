package main

import (
	"fmt"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(name string) error {
	return os.WriteFile(name, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error ocurred: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")
	return deck(s)
}

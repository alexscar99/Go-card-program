package main

import "fmt"

type deck []string

func newDeck() deck {
	// initialize empty var of type deck
	cards := deck{}

	// two string slices, one for suits and one for values
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// iteration to append each card, `_` is used as placeholder for the index vars not needed
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	// return var of type deck with all cards
	return cards
}

// receiver of type deck --> any var of type deck has access to print func
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

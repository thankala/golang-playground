package main

import "fmt"

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	card := "Nothing"
	figures := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}
	shapes := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	for _, f := range figures {
		for _, s := range shapes {
			card = f + " of " + s
			cards = append(cards, card)
		}
	}
	return cards
}

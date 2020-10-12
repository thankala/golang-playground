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
	figures := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}
	shapes := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	for _, figure := range figures {
		for _, shape := range shapes {
			cards = append(cards, figure+" of "+shape)
		}
	}
	return cards
}

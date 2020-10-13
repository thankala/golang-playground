package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

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

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(deckName string) {
	error := ioutil.WriteFile(deckName+".txt", []byte(d.toString()), 0644)
	if error != nil {
		log.Fatal(error)
	}
}

func readFromFile(deckName string) deck {
	content, error := ioutil.ReadFile(deckName + ".txt")
	if error != nil {
		log.Fatal(error)
	}
	return fromString(content)
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func fromString(content []byte) deck {
	return strings.Split(string(content), "\n")
}

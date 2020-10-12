package main

import "fmt"

func main() {
	deck := newDeck()
	for _, d := range deck {
		fmt.Println(d)

	}
	fmt.Println(len(deck))
}

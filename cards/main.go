package main

func main() {
	deck := readFromFile("mydeck")
	// fmt.Println(len(deck))

	// deck := newDeck()

	// deck.saveToFile("mydeck")
	// deck.shuffle()

	deck.print()

}

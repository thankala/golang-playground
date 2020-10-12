package main

func main() {
	// deck := newDeck()
	// for _, d := range deck {
	// 	fmt.Println(d)

	// }
	// fmt.Println(len(deck))

	deck := newDeck()

	deck.saveToFile("mydeck")

}

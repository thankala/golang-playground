package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf746",
		"white": "#ffffff",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	delete(colors, "white")
	fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for _, hex := range c {
		fmt.Println(hex)
	}
}

package main

import "fmt"

func main() {
	// var colors1 map[string]string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#74556",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

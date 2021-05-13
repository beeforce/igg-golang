package main

import (
	"fmt"
	"github.com/nitipatl/go-deck"
)

// import (
// 	deck "github.com/nitipatl/go-deck"
// )

func main() {
	cards := deck.NewDeck()
	// hand, remainingCards := deck.Deal(cards, 5)
	// hand.Print()
	// remainingCards.Print()
	// cards.Print()
	if err := cards.SaveToFile("deck.txt"); err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(cards.NewDeckFromFile("deck.txt"))
}
package main

import (
	"fmt"
	"main/deck"
)

func main() {
	// var card string = "10"
	// card:=10

	card := deck.NewDeck()
	// card.print()
	// card.deal()
	// hand1, hand2 := deal(card,5)

	// hand1.print()
	// fmt.Println("     ")
	// hand2.print()

	// fmt.Println(hand2.toString())
	card.SaveToFile("myhand.txt")
	fmt.Println("Haa bhai chalra hai-1")
	ans := deck.NewDeckFromFile("myhand.txt")
	fmt.Println("Haa bhai chalra hai-2")
	// fmt.Println(ans)
	ans.Print()
	ans.Shuffle()
	fmt.Println("<------------------------->")
	ans.Print()
}

func newCard() string {
	return "Haa bhai Chalra hai"
}

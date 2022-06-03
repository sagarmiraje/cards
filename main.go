package main

import "strings"

func main() {
	cards :=newDeck()
	cards.saveToFile("my_cards")
	myCards,err := readFromFile("my_cards")
	if nil == err{
		newCard := string(myCards)
		cards = strings.Split(newCard,",")
		//cards.print()
	}
	cards.shuffleDeck()
	cards.print()
}
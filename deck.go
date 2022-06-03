package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"math/rand"
	"time"
)
 
type deck []string

// function to print deck
func (p deck) print() {
	for i, card := range p{
		fmt.Println(i, card)
	}
}

//function to create new deck
func newDeck() deck {
	cards := deck{}
	cardSuits :=[]string{"Spades","Diamond","Heart","Clubs"}
	cardValues :=[]string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Queen","King","Joker"}
	for _,value :=range cardValues{
		for _,suit :=range cardSuits {
			cards = append(cards,value+" of "+suit)
		}
	}
	return cards
}
//functiont to create deal
func (d deck)deal(handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}

// function to convert array into single string
func (d deck) toString() string {
	return strings.Join([]string(d),",")
}

//function to save dec in file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName,[]byte(d.toString()),0666)
}

//function to read from file
func readFromFile(fileName string) ([]byte,error) {
	return ioutil.ReadFile(fileName)
}

func swap_values(a *string, b *string) {
	var tmp = *a
	*a = *b
	*b = tmp
  }

func (p deck) shuffleDeck() {
	for _, card := range p {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		random := r.Intn(len(p))
		swap_values(&p[random],&card)
	}
}
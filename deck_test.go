package main

import (
	"testing"
)

func TestNewDeck(t *testing.T){
	d := newDeck()
	if 52 == len(d) {
		t.Errorf("Expected length is 52 but we got: %d",len(d))
	} 
}

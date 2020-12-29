package main
import "testing"

func TestNewDeck(t *Testing.T){

  d:=newDeck()
  if(len(d)!=16) {
    t.Errorf("Expected deck length of 16, got length %v ", len(d))
  }

  if(d[0]!="Ace of Spades") {
t.Errorf("Expected first card Ace of Spade, got %v", d[0])

  }


    if(d[len(d)-1]!="Four of Clubs") {
  t.Errorf("Expected last card Four of Clubs, got %v", d[len(d)-1])

    }
}

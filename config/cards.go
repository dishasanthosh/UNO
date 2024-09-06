package config

import (
	"math/rand"
	"time"
)

//Card -> Single UNO card
type Card struct {
	Color string //red, green, blue, yellow, wild
	Value string // 0,1,2,3,4,5,6,7,8,9, skip, reverse, draw 2,wild, wild draw 4
}

//Deck -> Set of UNO card
type Deck struct {
	Cards []Card
}

//Initialize full deck of UNO cards(108 cards)
func NewDeck() *Deck {
	color := []string{"Blue", "Green", "Yellow", "Red"}
	value := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Skip", "Reverse", "Draw Two"}

	var cards []Card

	//add cards of all values to each color
	for _, color := range color {
		for _, value := range value {
			cards = append(cards, Card{Color: color, Value: value})
		}
	}

	//add wild cards
	for i := 0; i < 4; i++ {
		cards = append(cards, Card{Color: "Wild", Value: "Wild"})
		cards = append(cards, Card{Color: "Wild", Value: "Wild Draw 4"})
	}

	rand.Seed(time.Now().UnixNano()) //to make sure unique shuffle each time
	rand.Shuffle(len(cards), func(i,j int){
		cards[i], cards[j] = cards[j], cards[i]
	})
	return &Deck{Cards: cards}

}

//Draw a card from the deck
func(d *Deck) DrawCard() Card{
	card := d.Cards[0]
	d.Cards = d.Cards[1:] //Remove the drawn card from the deck
	return card
}

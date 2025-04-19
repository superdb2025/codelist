package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card struct represents a single card with suit and rank
type Card struct {
	suit string
	rank string
}

// NewCard creates a new card
func NewCard(suit, rank string) Card {
	return Card{suit: suit, rank: rank}
}

// String returns a string representation of the card (e.g., "Ace of Spades")
func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.rank, c.suit)
}

// Deck struct represents a deck of cards
type Deck struct {
	cards []Card
}

// NewDeck creates a new deck of 52 cards
func NewDeck() *Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	deck := []Card{}

	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, NewCard(suit, rank))
		}
	}

	return &Deck{cards: deck}
}

// Shuffle shuffles the cards in the deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d.cards {
		j := rand.Intn(len(d.cards))
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// DealCard removes the top card from the deck and returns it
func (d *Deck) DealCard() Card {
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card
}

// Player struct represents a player with a name and a card
type Player struct {
	name string
	card Card
}

// NewPlayer creates a new player
func NewPlayer(name string) Player {
	return Player{name: name}
}

// SetCard assigns a card to the player
func (p *Player) SetCard(card Card) {
	p.card = card
}

// ShowCard displays the player's card
func (p Player) ShowCard() {
	fmt.Printf("%s has: %s\n", p.name, p.card.String())
}

// GetCardValue returns a numerical value of the player's card for comparison
func (p Player) GetCardValue() int {
	rankValues := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
		"Jack": 11, "Queen": 12, "King": 13, "Ace": 14,
	}

	return rankValues[p.card.rank]
}

// Main function to run the game
func main() {
	// Create and shuffle the deck
	deck := NewDeck()
	deck.Shuffle()

	// Create two players
	player1 := NewPlayer("Player 1")
	player2 := NewPlayer("Player 2")

	// Deal one card to each player
	player1.SetCard(deck.DealCard())
	player2.SetCard(deck.DealCard())

	// Display the players' cards
	player1.ShowCard()
	player2.ShowCard()

	// Compare the cards and determine the winner
	if player1.GetCardValue() > player2.GetCardValue() {
		fmt.Println("Player 1 wins!")
	} else if player1.GetCardValue() < player2.GetCardValue() {
		fmt.Println("Player 2 wins!")
	} else {
		fmt.Println("It's a tie!")
	}
}

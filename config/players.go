package config

type Player struct {
    ID   string   // Unique player ID
    Hand []Card   // The player's hand of cards
}

// Game represents an ongoing UNO game.
type Game struct {
    Players       []*Player
    Deck          *Deck
    DiscardPile   []Card
    CurrentPlayer int // Index of the current player's turn
    Direction     int // 1 for clockwise, -1 for counter-clockwise
}

func NewGame(playerIDs []string) *Game {
    deck := NewDeck()
    players := make([]*Player, len(playerIDs))

    // Initialize each player with 7 cards
    for i, id := range playerIDs {
        players[i] = &Player{
            ID:   id,
            Hand: drawInitialHand(deck),
        }
    }

    // Initialize the game with the deck, players, and a starting discard pile.
    return &Game{
        Players:     players,
        Deck:        deck,
        DiscardPile: []Card{deck.DrawCard()}, // Starting card on the discard pile
        Direction:   1,                       // Clockwise
    }
}

// Function to draw 7 cards for each player at the start of the game.
func drawInitialHand(deck *Deck) []Card {
    var hand []Card
    for i := 0; i < 7; i++ {
        hand = append(hand, deck.DrawCard())
    }
    return hand
}
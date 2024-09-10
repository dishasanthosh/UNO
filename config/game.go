package config
import (
	"fmt"
)
// Helper function to remove a card from a player's hand
func removeCardFromHand(hand *[]Card, card Card) {
    for i, c := range *hand {
        if c.Color == card.Color && c.Value == card.Value {
            // Remove the card from the player's hand by slicing it out
            *hand = append((*hand)[:i], (*hand)[i+1:]...)
            break
        }
    }
}

// Handle a player's turn and apply the game logic based on the card they play.
func (g *Game) PlayCard(playerID string, playedCard Card) error {
    // Ensure it's the player's turn
    if g.Players[g.CurrentPlayer].ID != playerID {
        return fmt.Errorf("not your turn")
    }

    // Ensure the card is playable (matches color or value with the top of the discard pile)
    topCard := g.DiscardPile[len(g.DiscardPile)-1]
    if !(playedCard.Color == topCard.Color || playedCard.Value == topCard.Value || playedCard.Color == "Wild") {
        return fmt.Errorf("card not playable")
    }

    // Remove the card from the player's hand and add it to the discard pile
    player := g.Players[g.CurrentPlayer]
    g.DiscardPile = append(g.DiscardPile, playedCard)
    removeCardFromHand(&player.Hand, playedCard)

    // Apply special card actions (Skip, Reverse, Draw Two, Wild)
    switch playedCard.Value {
    case "Skip":
        g.CurrentPlayer = (g.CurrentPlayer + g.Direction + len(g.Players)) % len(g.Players)
    case "Reverse":
        g.Direction *= -1
    case "Draw Two":
        nextPlayer := (g.CurrentPlayer + g.Direction + len(g.Players)) % len(g.Players)
        g.Players[nextPlayer].Hand = append(g.Players[nextPlayer].Hand, g.Deck.DrawCard(), g.Deck.DrawCard())
    case "Wild", "Wild Draw Four":
        // Handle wild card actions (color choice, drawing cards)
    }

    // Advance to the next player's turn
    g.CurrentPlayer = (g.CurrentPlayer + g.Direction + len(g.Players)) % len(g.Players)
    return nil
}

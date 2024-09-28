package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen":
		return 10
	case "jack":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var sumOfCards int = ParseCard(card1) + ParseCard(card2)
	var dealerCardValue int = ParseCard(dealerCard)
	var action string

	switch {
	case card1 == "ace" && card2 == "ace":
		action = "P"
	case (sumOfCards == 21) && dealerCardValue < 10:
		action = "W"
	case (sumOfCards == 21) && dealerCardValue >= 10:
		action = "S"
	case sumOfCards >= 17 && sumOfCards <= 20:
		action = "S"
	case (12 <= sumOfCards && sumOfCards <= 16) && dealerCardValue < 7:
		action = "S"
	case (12 <= sumOfCards && sumOfCards <= 16) && dealerCardValue >= 7:
		action = "H"
	case sumOfCards <= 11:
		action = "H"
	}

	return action
}

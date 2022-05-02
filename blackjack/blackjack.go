package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var s int

	switch card {
	case "ace":
		s = 11
	case "ten", "jack", "queen", "king":
		s = 10
	case "nine":
		s = 9
	case "eight":
		s = 8
	case "seven":
		s = 7
	case "six":
		s = 6
	case "five":
		s = 5
	case "four":
		s = 4
	case "three":
		s = 3
	case "two":
		s = 2
	default:
		s = 0
	}

	return s
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	v := ParseCard(card1) + ParseCard(card2)
	d := ParseCard(dealerCard)

	var s string

	switch {
	case v == 22:
		s = "P"

	case v == 21:
		switch d {
		case 11, 10:
			s = "S"
		default:
			s = "W"
		}

	case v >= 17 && v <= 20:
		s = "S"

	case v >= 12 && v <= 16:
		if d >= 7 {
			s = "H"
		} else {
			s = "S"
		}

	case v <= 11:
		s = "H"
	}

	return s
}

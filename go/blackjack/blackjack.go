package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardNum := 0
	switch {
	case card == "two":
		cardNum = 2
	case card == "three":
		cardNum = 3
	case card == "four":
		cardNum = 4
	case card == "five":
		cardNum = 5
	case card == "six":
		cardNum = 6
	case card == "seven":
		cardNum = 7
	case card == "eight":
		cardNum = 8
	case card == "nine":
		cardNum = 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		cardNum = 10
	case card == "ace":
		cardNum = 11
	default:
		cardNum = 0
	}
	return cardNum
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if ParseCard(card1)+ParseCard(card2) == 22 {
		return "P"
	} else if ParseCard(card1)+ParseCard(card2) == 21 && (ParseCard(dealerCard) != 11 && ParseCard(dealerCard) != 10) {
		return "W"
	} else if ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) == 11 {
		return "S"
	} else if ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) == 10 {
		return "S"
	} else if ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20 {
		return "S"
	} else if ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(dealerCard) < 7 {
		return "S"
	} else if ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(dealerCard) >= 7 {
		return "H"
	} else {
		return "H"
	}
}

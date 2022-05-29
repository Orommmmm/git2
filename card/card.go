package card

import (
	"orom/types"
)

func Vvod() types.Card {
	card := types.Card{
		ID:       0,
		PAN:      "5050 xxxx xxxx 6969",
		Currency: "TJS",
		Color:    "Gold",
		Name:     "Boozoorg",
		Balance:  20000,
		Active:   true,
	}
	return Withdraw(&card)
}

func Withdraw(card *types.Card) types.Card {
	if !card.Active {
		return *card
	} else {
		card.Active = false
	}
	return *card
}

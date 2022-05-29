package card

import (
	"fmt"
	"orom/types"
)

func Vvod() types.Card {
	card := types.Card{
		ID:       0,
		PAN:      "5050 xxxx xxxx 0001",
		Currency: "TJS",
		Color:    "Platinum",
		Name:     "Orom",
		Balance:  60000,
		Active:   true,
	}
	fmt.Println(Withdraw(&card))
}

func Withdraw(card *types.Card) types.Card {
	if !card.Active {
		return *card
	} else {
		card.Active = false
	}
	return *card
}

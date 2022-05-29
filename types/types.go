package types

type Money int64

type PAN string

type Currency string

const (
	TJS  Currency = "TJS"
	RUB  Currency = "RUB"
	USD  Currency = "USD"
	Euro Currency = "EUR"
)

type Color string

const (
	Silver Color = "Platinum"
	Gold   Color = "Gold"
)

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    Color
	Name     string
	Active   bool
}

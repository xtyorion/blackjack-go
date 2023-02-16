package blackjack

import "fmt"

type Card struct {
	Suit  string
	Rank  string
	Value int
}

var Suits = []string{"Clubs", "Diamonds", "Hearts", "Spades"}
var Ranks = map[string]int{
	"Ace":   10,
	"King":  10,
	"Queen": 10,
	"Jack":  10,
	"10":    10,
	"9":     9,
	"8":     8,
	"7":     7,
	"6":     6,
	"5":     5,
	"4":     4,
	"3":     3,
	"2":     2,
}

func (c *Card) ToString() string {
	return c.Suit + " of " + c.Rank + " " + fmt.Sprint(c.Value)
}

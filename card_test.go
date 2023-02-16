package blackjack

import (
	"testing"
)

func TestCard(t *testing.T) {
	t.Run("Card name getter", func(t *testing.T) {
		card := Card{"Ace", "Clubs", 10}
		want := "Ace of Clubs 10"
		got := card.ToString()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

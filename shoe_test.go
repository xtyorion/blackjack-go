package blackjack

import (
	"testing"
)

func TestShoe(t *testing.T) {
	t.Run("Initialization of Shoe", func(t *testing.T) {
		shoe := Shoe{}
		shoe.Initialize()

		want := 1
		got := shoe.NumberOfDecks

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

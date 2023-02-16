package blackjack

import (
	"fmt"
	"math/rand"
	"time"
)

type Shoe struct {
	Cards          []Card
	NumberOfDecks  int
	CurrentCardIdx int
}

func (s *Shoe) Initialize() {
	s.NumberOfDecks = 1
	s.repopulate()
	s.printDeck()
}

func (s *Shoe) Draw() Card {
	if s.CurrentCardIdx >= len(s.Cards) {
		fmt.Print("Throw Error, empty shoe")
	}

	retVal := s.Cards[s.CurrentCardIdx]
	s.NextCard()

	return retVal
}

func (s *Shoe) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s.Cards), func(i, j int) {
		s.Cards[i], s.Cards[j] = s.Cards[j], s.Cards[i]
	})
}

func (s *Shoe) NextCard() {
	s.CurrentCardIdx += 1
}

// Returns an array of Card that already been dealt
func (s *Shoe) DealtCards() []Card {
	//ranges from 0 to current card index
	return s.Cards[0:s.CurrentCardIdx]
}

// Returns an array of Card that hasn't been dealt
func (s *Shoe) UndealtCards() []Card {
	//ranges from current card index to
	return s.Cards[s.CurrentCardIdx : len(s.Cards)-1]
}

func (s *Shoe) NumberOfCards() int {
	return len(s.Cards) - s.CurrentCardIdx
}

func (s *Shoe) repopulate() {
	s.Cards = nil

	//for each deck in number of decks
	for x := 0; x < s.NumberOfDecks; x++ {
		for _, suitVal := range Suits {
			for rankKey, rankVal := range Ranks {
				s.Cards = append(s.Cards, Card{suitVal, rankKey, rankVal})
			}
		}
	}
	s.Shuffle()
}

func (s *Shoe) printDeck() {
	for _, val := range s.Cards {
		fmt.Println(val.ToString())
	}
}

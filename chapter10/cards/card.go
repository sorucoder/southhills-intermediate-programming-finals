package cards

import (
	"errors"
	"fmt"
)

var (
	ErrOutOfCards     error = errors.New("out of cards")
	ErrInvalidDealing error = errors.New("invalid dealing")
)

type Suit int

const (
	SPADES Suit = iota
	CLUBS
	HEARTS
	DIAMONDS
)

func (suit Suit) String() string {
	switch suit {
	case SPADES:
		return "Spades"
	case CLUBS:
		return "Clubs"
	case HEARTS:
		return "Hearts"
	case DIAMONDS:
		return "Diamonds"
	default:
		return ""
	}
}

type Rank int

const (
	ACE Rank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

func (rank Rank) String() string {
	switch rank {
	case ACE:
		return "Ace"
	case TWO:
		return "2"
	case THREE:
		return "3"
	case FOUR:
		return "4"
	case FIVE:
		return "5"
	case SIX:
		return "6"
	case SEVEN:
		return "7"
	case EIGHT:
		return "8"
	case NINE:
		return "9"
	case TEN:
		return "10"
	case JACK:
		return "Jack"
	case QUEEN:
		return "Queen"
	case KING:
		return "King"
	default:
		return ""
	}
}

type Card struct {
	rank Rank
	suit Suit
}

func MakeCard(rank Rank, suit Suit) *Card {
	return &Card{rank, suit}
}

func (card *Card) Rank() Rank {
	return card.rank
}

func (card *Card) SetRank(rank Rank) {
	card.rank = rank
}

func (card *Card) Suit() Suit {
	return card.suit
}

func (card *Card) SetSuit(suit Suit) {
	card.suit = suit
}

func (card *Card) Equals(other *Card) bool {
	if card == other {
		return true
	}
	return card.rank == other.rank && card.suit == other.suit
}

func (card *Card) String() string {
	return fmt.Sprintf("%s of %s", card.rank, card.suit)
}

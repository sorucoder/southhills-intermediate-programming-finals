package main

import (
	"fmt"
	"math/rand"

	"github.com/sorucoder/southhills-it100-finals/chapter10/cards"
)

func drawCard(hand *[]*cards.Card) *cards.Card {
	if len(*hand) > 0 {
		card := (*hand)[0]
		*hand = (*hand)[1:]
		return card
	}
	return nil
}

func declareWar(hand *[]*cards.Card) []*cards.Card {
	var count int
	if len(*hand) < 4 {
		count = len(*hand)
	} else {
		count = 4
	}
	cards := (*hand)[:count-1]
	*hand = (*hand)[count-1:]
	return cards
}

func main() {
	fmt.Println("Shuffling deck...")
	deck := make([]*cards.Card, 52)
	for index, value := range rand.Perm(52) {
		deck[index] = cards.MakeCard(cards.Rank(value%13), cards.Suit(value/13))
	}

	humanHand, computerHand := deck[0:26], deck[26:]
	var humanCard, computerCard *cards.Card
	for len(humanHand) > 0 && len(computerHand) > 0 {
		humanCard, computerCard = drawCard(&humanHand), drawCard(&computerHand)
		fmt.Printf("You drew the %v!\n", humanCard)
		fmt.Printf("They drew the %v!\n", computerCard)
		if humanCard.Rank() > computerCard.Rank() {
			humanHand = append(humanHand, humanCard, computerCard)
			fmt.Println("You won this round.")
		} else if humanCard.Rank() < computerCard.Rank() {
			computerHand = append(computerHand, humanCard, computerCard)
			fmt.Println("They won this round.")
		} else {
			warCards := make([]*cards.Card, 0)
			for len(humanHand) > 0 && len(computerHand) > 0 && humanCard.Rank() == computerCard.Rank() {
				fmt.Println("I")
				fmt.Println("Declare")
				fmt.Println("WAR!")
				warCards = append(warCards, declareWar(&humanHand)...)
				warCards = append(warCards, declareWar(&computerHand)...)

				humanCard, computerCard = drawCard(&humanHand), drawCard(&computerHand)
				fmt.Printf("You drew the %v!\n", humanCard)
				fmt.Printf("They drew the %v!\n", computerCard)
				warCards = append(warCards, humanCard, computerCard)
			}

			if len(computerHand) == 0 || humanCard.Rank() > computerCard.Rank() {
				humanHand = append(humanHand, warCards...)
			} else if len(humanHand) == 0 || humanCard.Rank() < computerCard.Rank() {
				computerHand = append(computerHand, warCards...)
			}
		}
		fmt.Printf("You: %d Them: %d\n", len(humanHand), len(computerHand))
	}

	if len(humanHand) > 0 {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}
}

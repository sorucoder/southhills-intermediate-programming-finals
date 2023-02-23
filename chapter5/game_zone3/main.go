package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

var help string = `PIG
=================================================================================================================
The rules of the game are as follows:
* At the beginning, a coin will be flipped to determine if the human or computer goes first.
* Each player will roll dice until the end of their turn, adding the sum of the two dice to their total score.
* If a player does not roll any 1s, that player's will add the sum of the two dice to their score, and continue.
* If a player rolls one 1, that player's turn is over, and they forfeit the score they earned for that turn.
* If a player rolls two 1s, that player's turn is over, and they forfeit their total score.
* Either player can choose to end their turn at any time.
* The first player to reach a total score of 100 wins.
`
var randomizer *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	fmt.Println(help)
	var humanTotalScore, computerTotalScore int
	var humanTurnScore, computerTurnScore int
	turn := randomizer.Int() % 2

	for humanTotalScore+humanTurnScore < 100 && computerTotalScore+computerTurnScore < 100 {
		if turn%2 == 0 {
			fmt.Println("\033[1m\033[32mYour turn!\033[0m")

			var roll bool
			if errAskToRoll := survey.AskOne(
				&survey.Confirm{
					Message: "Would you like to roll the dice?",
					Help:    help,
					Default: true,
				},
				&roll,
			); errAskToRoll != nil {
				panic(errAskToRoll)
			}

			randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
			if roll {
				firstRoll := randomizer.Int()%6 + 1
				secondRoll := randomizer.Int()%6 + 1
				fmt.Printf("\033[32mYou rolled a %d and a %d.\033[0m\n", firstRoll, secondRoll)
				if firstRoll == 1 || secondRoll == 1 {
					if firstRoll == 1 && secondRoll == 1 {
						humanTotalScore = 0
					}
					humanTurnScore = 0
					turn++
				} else {
					humanTurnScore += firstRoll + secondRoll
				}
			} else {
				fmt.Println("\033[1m\033[32mYou decided not to roll.\033[0m")
				humanTotalScore += humanTurnScore
				humanTurnScore = 0
				turn++
			}
		} else {
			fmt.Println("\033[1m\033[31mTheir turn!\033[0m")

			roll := randomizer.Int()%100 >= 50

			if roll {
				firstRoll := randomizer.Int()%6 + 1
				secondRoll := randomizer.Int()%6 + 1
				fmt.Printf("\033[31mThey rolled a %d and a %d.\033[0m\n", firstRoll, secondRoll)
				if firstRoll == 1 || secondRoll == 1 {
					if firstRoll == 1 && secondRoll == 1 {
						computerTotalScore = 0
					}
					computerTurnScore = 0
					turn++
				} else {
					computerTurnScore += firstRoll + secondRoll
				}
			} else {
				fmt.Println("\033[31mThey decided not to roll.\033[0m")
				computerTotalScore += computerTurnScore
				computerTurnScore = 0
				turn++
			}
		}
		fmt.Printf("\033[1m\033[33mYou: %d Them: %d\033[0m\n", humanTotalScore+humanTurnScore, computerTotalScore+computerTurnScore)
	}
	humanTotalScore += humanTurnScore
	computerTotalScore += humanTurnScore

	if humanTotalScore > computerTotalScore {
		fmt.Println(`You win!`)
	} else {
		fmt.Println(`You lose!`)
	}
}

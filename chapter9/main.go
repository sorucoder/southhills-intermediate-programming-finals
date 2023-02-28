package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

var errInvalidPrompt error = errors.New("invalid prompt")

func validateAmount(answer any) error {
	if answerValue, ok := answer.(string); ok {
		_, errParse := strconv.ParseFloat(answerValue, 64)
		return errParse
	}
	return errInvalidPrompt
}

func calculateBills(amount, denomination float64) (count int, remainder float64) {
	count = int(amount / denomination)
	remainder = amount - float64(count)*denomination
	return
}

var denominations map[float64]string = map[float64]string{
	100.00: "$100 Bills",
	50.00:  "$50 Bills",
	20.00:  "$20 Bills",
	10.00:  "$10 Bills",
	5.00:   "$5 Bills",
	1.00:   "$1 Bills",
	0.25:   "Quarters",
	0.10:   "Dimes",
	0.05:   "Nickels",
	0.01:   "Pennies",
}

func main() {
	var totalAmount float64
	for {
		var amountAnswer string
		if errAskDollarAmount := survey.AskOne(
			&survey.Input{
				Message: "Please enter a dollar amount:",
				Default: "0",
				Help:    "Enter an amount in USD, without units or currency symbols. Enter 0 when finished.",
			},
			&amountAnswer,
			survey.WithValidator(survey.ComposeValidators(survey.Required, validateAmount)),
		); errAskDollarAmount != nil {
			panic(errAskDollarAmount)
		}

		amount, _ := strconv.ParseFloat(amountAnswer, 64)
		if amount == 0 {
			break
		}

		totalAmount += amount
		fmt.Printf("Total Amount: $%.2f\n", totalAmount)
	}
	totalAmount = math.Round(totalAmount*100) / 100

	fmt.Printf("$%.2f breaks down as follows:\n", totalAmount)
	for denominationAmount, denominationName := range denominations {
		var bills int
		bills, totalAmount = calculateBills(totalAmount, denominationAmount)

		if bills != 0 {
			fmt.Printf("%d %s\n", bills, denominationName)
		}

		if totalAmount <= 0 {
			break
		}
	}
}

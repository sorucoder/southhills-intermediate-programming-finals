package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
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

var denominationAmounts []float64 = []float64{
	100.00,
	50.00,
	20.00,
	10.00,
	5.00,
	1.00,
	0.25,
	0.10,
	0.05,
	0.01,
}

var denominationAmountsSorter sort.Interface = sort.Reverse(sort.Float64Slice(denominationAmounts))

var denominationSingularNames map[float64]string = map[float64]string{
	100.00: "$100 Bill",
	50.00:  "$50 Bill",
	20.00:  "$20 Bill",
	10.00:  "$10 Bill",
	5.00:   "$5 Bill",
	1.00:   "$1 Bill",
	0.25:   "Quarter",
	0.10:   "Dime",
	0.05:   "Nickel",
	0.01:   "Penny",
}

var denominationPluralNames map[float64]string = map[float64]string{
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

func calculateConsecutiveAmounts(amount float64) map[float64]int {
	consecutiveAmounts := make(map[float64]int)

	if !sort.IsSorted(denominationAmountsSorter) {
		sort.Sort(denominationAmountsSorter)
	}

	for _, denominationAmount := range denominationAmounts {
		count := int(amount / denominationAmount)
		if count != 0 {
			consecutiveAmounts[denominationAmount] = count
			amount -= float64(count) * denominationAmount
			if amount <= 0 {
				break
			}
		}
	}
	return consecutiveAmounts
}

func main() {
	var totalAmount float64
	for {
		var amountAnswer string
		if errAskDollarAmount := survey.AskOne(
			&survey.Input{
				Message: "Please enter a dollar amount:",
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
	consecutiveAmounts := calculateConsecutiveAmounts(totalAmount)
	for denominationValue, count := range consecutiveAmounts {
		var denominationName string
		if count == 1 {
			denominationName = denominationSingularNames[denominationValue]
		} else {
			denominationName = denominationPluralNames[denominationValue]
		}
		fmt.Printf("%d %s\n", count, denominationName)
	}
}

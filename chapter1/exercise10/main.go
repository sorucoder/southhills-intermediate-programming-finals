package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

var errInvalidPrompt error = errors.New("invalid prompt")

func validateInteger(answer any) error {
	if answerValue, ok := answer.(string); ok {
		_, errParse := strconv.ParseInt(answerValue, 10, 64)
		return errParse
	}
	return errInvalidPrompt
}

func transformInteger(answer any) any {
	if answerValue, ok := answer.(string); ok {
		integerValue, errParse := strconv.ParseInt(answerValue, 10, 64)
		if errParse != nil {
			panic(errParse)
		}
		return integerValue
	}
	panic(errInvalidPrompt)
}

func validateFloat(answer any) error {
	if answerValue, ok := answer.(string); ok {
		_, errParse := strconv.ParseFloat(answerValue, 64)
		return errParse
	}
	return errInvalidPrompt
}

func transformFloat(answer any) any {
	if answerValue, ok := answer.(string); ok {
		amount, errParse := strconv.ParseFloat(answerValue, 64)
		if errParse != nil {
			panic(errParse)
		}
		return amount
	}
	panic(errInvalidPrompt)
}

var salespersonIncomeSurvey []*survey.Question = []*survey.Question{
	{
		Name: "salary",
		Prompt: &survey.Input{
			Message: "Please enter your base salary:",
			Help:    "You can enter an integer or decimal value, without the currency symbol.",
		},
		Validate:  survey.ComposeValidators(survey.Required, validateFloat),
		Transform: transformFloat,
	},
	{
		Name: "sales",
		Prompt: &survey.Input{
			Message: "Please enter your total sales:",
			Help:    "You can enter an integer value.",
		},
		Validate:  survey.ComposeValidators(survey.Required, validateInteger),
		Transform: transformInteger,
	},
	{
		Name: "commission",
		Prompt: &survey.Input{
			Message: "Please enter your commission rate:",
			Help:    "You can enter an integer or decimal value, without the currency symbol.",
		},
		Validate:  survey.ComposeValidators(survey.Required, validateFloat),
		Transform: transformFloat,
	},
}

type salespersonIncomeAnswers struct {
	Salary     float64 `survey:"salary"`
	Sales      int64   `survey:"sales"`
	Commission float64 `survey:"commissions"`
}

func main() {
	var answers salespersonIncomeAnswers
	if errAsk := survey.Ask(salespersonIncomeSurvey, &answers); errAsk != nil {
		panic(errAsk)
	}

	fmt.Printf("Your Pay: $%.2f\n", answers.Salary+float64(answers.Sales)*answers.Commission)
}

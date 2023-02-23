package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sorucoder/southhills-it100-finals/chapter1/exercise9/exchangerate"
)

var errInvalidPrompt error = errors.New("invalid prompt")

func suggestCurrency(partial string) []string {
	partialLower, partialUpper := strings.ToLower(partial), strings.ToUpper(partial)

	suggestions := make([]string, 0)

	// Try to suggest codes first. If it suggestions are found, return them.
	if len(partial) <= 3 {
		for _, currency := range exchangerate.SupportedCurrencies {
			if strings.HasPrefix(strings.ToUpper(currency.Code), partialUpper) {
				suggestions = append(suggestions, currency.String())
			}
		}
		if len(suggestions) > 0 {
			return suggestions
		}
	}

	// Otherwise, go for names.
	for _, currency := range exchangerate.SupportedCurrencies {
		if strings.Contains(strings.ToLower(currency.Name), partialLower) {
			suggestions = append(suggestions, currency.String())
		}
	}
	return suggestions
}

func validateCurrency(answer any) error {
	if answerValue, ok := answer.(string); ok {
		for _, currency := range exchangerate.SupportedCurrencies {
			if answerValue == currency.String() {
				return nil
			}
		}
		return exchangerate.ErrUnsupportedCurrency
	}
	return errInvalidPrompt
}

func transformCurrency(answer any) any {
	if answerValue, ok := answer.(string); ok {
		for _, currency := range exchangerate.SupportedCurrencies {
			if answerValue == currency.String() {
				return currency
			}
		}
	}
	panic(errInvalidPrompt)
}

func validateAmount(answer any) error {
	if answerValue, ok := answer.(string); ok {
		_, errParse := strconv.ParseFloat(answerValue, 64)
		return errParse
	}
	return errInvalidPrompt
}

func transformAmount(answer any) any {
	if answerValue, ok := answer.(string); ok {
		amount, errParse := strconv.ParseFloat(answerValue, 64)
		if errParse != nil {
			panic(errParse)
		}
		return amount
	}
	panic(errInvalidPrompt)
}

var convertCurrencySurvey []*survey.Question = []*survey.Question{
	{
		Name: "fromCurrency",
		Prompt: &survey.Input{
			Message: "Please enter the currency you wish to convert from:",
			Help:    "You can enter either an ISO 4217 currency code or a currency name, and use Tab for completions.",
			Suggest: suggestCurrency,
		},
		Validate:  survey.ComposeValidators(survey.Required, validateCurrency),
		Transform: transformCurrency,
	},
	{
		Name: "toCurrency",
		Prompt: &survey.Input{
			Message: "Please enter the currency you wish to convert to:",
			Help:    "You can enter either an ISO 4217 currency code or a currency name, and use Tab for completions.",
			Suggest: suggestCurrency,
		},
		Validate:  survey.ComposeValidators(survey.Required, validateCurrency),
		Transform: transformCurrency,
	},
	{
		Name: "amount",
		Prompt: &survey.Input{
			Message: "Please enter the amount you wish to convert:",
			Help:    "You can enter an integer or decimal value, without the currency symbol.",
		},
		Validate:  survey.ComposeValidators(survey.Required, validateAmount),
		Transform: transformAmount,
	},
}

type convertCurrencyAnswers struct {
	FromCurrency *exchangerate.Currency `survey:"fromCurrency"`
	ToCurrency   *exchangerate.Currency `survey:"toCurrency"`
	Amount       float64                `survey:"amount"`
}

func main() {
	var answers convertCurrencyAnswers
	errAsk := survey.Ask(convertCurrencySurvey, &answers)
	if errAsk != nil {
		panic(errAsk)
	}
	toAmount := exchangerate.Convert(answers.FromCurrency, answers.ToCurrency, answers.Amount)
	fmt.Printf("%.4f %s = %.4f %s\n", answers.Amount, answers.FromCurrency.Code, toAmount, answers.ToCurrency.Code)
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sorucoder/southhills-it100-finals/chapter2/arniesappliances"
)

var errInvalidPrompt error = errors.New("invalid prompt")

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

func main() {
	var refrigerators []*arniesappliances.RefrigeratorRecord
	if jsonFile, errOpen := os.Open("records.json"); errOpen == nil {
		jsonDecoder := json.NewDecoder(jsonFile)
		if errDecoder := jsonDecoder.Decode(&refrigerators); errDecoder != nil {
			panic(errDecoder)
		}
		jsonFile.Close()
	} else {
		refrigerators = make([]*arniesappliances.RefrigeratorRecord, 0)
	}

	var done bool
	for !done {
		var modelAnswer string
		if errAskRefrigeratorModel := survey.AskOne(
			&survey.Input{
				Message: "Please enter the model name of the refrigerator:",
				Default: "XXX",
				Help:    `You can enter the model name of this refrigerator; or leave it blank or enter "XXX" to finish`,
			},
			&modelAnswer,
		); errAskRefrigeratorModel != nil {
			panic(errAskRefrigeratorModel)
		}

		if modelAnswer != "XXX" {
			var interiorDimensionsAnswer struct {
				Width  float64 `survey:"width"`
				Height float64 `survey:"height"`
				Depth  float64 `survey:"depth"`
			}
			if errAskRefrigeratorDimensions := survey.Ask(
				[]*survey.Question{
					{
						Name: "width",
						Prompt: &survey.Input{
							Message: "Please enter the interior width of the refrigerator in cubic inches:",
							Help:    "You can enter an integer or decimal value, without the unit.",
						},
						Validate:  validateFloat,
						Transform: transformFloat,
					},
					{
						Name: "height",
						Prompt: &survey.Input{
							Message: "Please enter the interior height of the refrigerator in cubic inches:",
							Help:    "You can enter an integer or decimal value, without the unit.",
						},
						Validate:  validateFloat,
						Transform: transformFloat,
					},
					{
						Name: "depth",
						Prompt: &survey.Input{
							Message: "Please enter the interior width of the refrigerator in cubic inches:",
							Help:    "You can enter an integer or decimal value, without the unit.",
						},
						Validate:  validateFloat,
						Transform: transformFloat,
					},
				},
				&interiorDimensionsAnswer,
			); errAskRefrigeratorDimensions != nil {
				panic(errAskRefrigeratorDimensions)
			}
			refrigerators = append(refrigerators, arniesappliances.NewRefrigeratorRecord(modelAnswer, interiorDimensionsAnswer.Width, interiorDimensionsAnswer.Height, interiorDimensionsAnswer.Depth))
		} else {
			done = true
		}
	}

	if jsonFile, errCreate := os.Create("records.json"); errCreate == nil {
		defer jsonFile.Close()
		jsonEncoder := json.NewEncoder(jsonFile)
		if errEncode := jsonEncoder.Encode(refrigerators); errEncode != nil {
			panic(errEncode)
		}
	}

	fmt.Println("End of job")
}

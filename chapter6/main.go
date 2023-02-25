package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sorucoder/southhills-it100-finals/chapter6/dailylifemagazine"
)

var (
	errInvalidPrompt        error = errors.New("invalid prompt")
	errInvalidAge           error = errors.New("invalid age")
	errInvalidGender        error = errors.New("invalid gender")
	errInvalidMaritalStatus error = errors.New("invalid marital status")
	errInvalidAnnualIncome  error = errors.New("invalid annual income")
)

func validateAge(answer any) error {
	if answerValue, ok := answer.(string); ok {
		age, errParse := strconv.ParseInt(answerValue, 10, 64)
		if errParse != nil {
			return errParse
		} else if age < 0 || age > 120 {
			return errInvalidAge
		}
		return nil
	}
	return errInvalidPrompt
}

func transformAge(answer any) any {
	if answerValue, ok := answer.(string); ok {
		age, errParse := strconv.ParseInt(answerValue, 10, 64)
		if errParse != nil {
			panic(errParse)
		} else if age < 0 || age > 120 {
			panic(errInvalidAge)
		}
		return int(age)
	}
	panic(errInvalidPrompt)
}

func validateGender(answer any) error {
	if option, ok := answer.(survey.OptionAnswer); ok {
		if option.Value == "Male" || option.Value == "Female" {
			return nil
		}
		return errInvalidAge
	}
	return errInvalidPrompt
}

func validateMaritalStatus(answer any) error {
	if option, ok := answer.(survey.OptionAnswer); ok {
		if option.Value == "Single" || option.Value == "Married" || option.Value == "Widowed" || option.Value == "Divorced" {
			return nil
		}
		return errInvalidMaritalStatus
	}
	return errInvalidPrompt
}

func validateAnnualIncome(answer any) error {
	if answerValue, ok := answer.(string); ok {
		annualIncome, errParse := strconv.ParseInt(answerValue, 10, 64)
		if errParse != nil {
			return errParse
		} else if annualIncome < 15000 || annualIncome > 1000000000000 {
			return errInvalidAnnualIncome
		}
		return nil
	}
	return errInvalidPrompt
}

func transformAnnualIncome(answer any) any {
	if answerValue, ok := answer.(string); ok {
		annualIncome, errParse := strconv.ParseInt(answerValue, 10, 64)
		if errParse != nil {
			panic(errParse)
		} else if annualIncome < 15000 || annualIncome > 1000000000000 {
			panic(errInvalidAge)
		}
		return annualIncome
	}
	panic(errInvalidPrompt)
}

var addReaderSurvey []*survey.Question = []*survey.Question{
	{
		Name: "age",
		Prompt: &survey.Input{
			Message: "Please enter the age of the reader:",
			Help:    `You can enter a positive integer in years old, without "years old".`,
		},
		Validate:  survey.ComposeValidators(survey.Required, validateAge),
		Transform: transformAge,
	},
	{
		Name: "gender",
		Prompt: &survey.Select{
			Message: "Please select the gender of the reader:",
			Options: []string{
				"Male",
				"Female",
			},
		},
		Validate: survey.ComposeValidators(survey.Required, validateGender),
	},
	{
		Name: "maritalStatus",
		Prompt: &survey.Select{
			Message: "Please select the marital status of the reader:",
			Options: []string{
				"Single",
				"Married",
				"Widowed",
				"Divorced",
			},
		},
		Validate: survey.ComposeValidators(survey.Required, validateMaritalStatus),
	},
	{
		Name: "annualIncome",
		Prompt: &survey.Input{
			Message: "Please enter the annual income of the reader:",
			Help:    `You can enter a positive integer in United States Dollars, without the unit.`,
		},
		Validate:  survey.ComposeValidators(survey.Required, validateAnnualIncome),
		Transform: transformAnnualIncome,
	},
}

func loadRecords() ([]*dailylifemagazine.Reader, error) {
	var readers []*dailylifemagazine.Reader
	if jsonFile, errOpen := os.Open("records.json"); errOpen == nil {
		jsonDecoder := json.NewDecoder(jsonFile)
		if errDecode := jsonDecoder.Decode(&readers); errDecode != nil {
			return nil, errDecode
		}
		jsonFile.Close()
	} else {
		readers = make([]*dailylifemagazine.Reader, 0)
	}
	return readers, nil
}

func saveRecords(readers []*dailylifemagazine.Reader) error {
	if jsonFile, errCreate := os.Create("records.json"); errCreate == nil {
		jsonEncoder := json.NewEncoder(jsonFile)
		if errEncode := jsonEncoder.Encode(readers); errEncode != nil {
			return errEncode
		}
		jsonFile.Close()
	} else {
		return errCreate
	}
	return nil
}

func main() {
	var menuOption string
	if errAskMenuOption := survey.AskOne(
		&survey.Select{
			Message: "Please select an option",
			Options: []string{
				"Add readers...",
				"View readers by age group...",
				"View readers by gender and age group...",
				"View readers by annual income...",
			},
		},
		&menuOption,
	); errAskMenuOption != nil {
		panic(errAskMenuOption)
	}

	switch menuOption {
	case "Add readers...":
		readers, errLoad := loadRecords()
		if errLoad != nil {
			panic(errLoad)
		}

		addingReaders := true
		for addingReaders {
			var answers struct {
				Age           int    `survey:"age"`
				Gender        string `survey:"gender"`
				MaritalStatus string `survey:"maritalStatus"`
				AnnualIncome  int64  `survey:"annualIncome"`
			}
			if errAskAddReaderSurvey := survey.Ask(addReaderSurvey, &answers); errAskAddReaderSurvey != nil {
				panic(errAskAddReaderSurvey)
			}

			var gender dailylifemagazine.Gender
			switch answers.Gender {
			case "Male":
				gender = dailylifemagazine.GENDER_MALE
			case "Female":
				gender = dailylifemagazine.GENDER_FEMALE
			default:
				panic(errInvalidGender)
			}

			var maritalStatus dailylifemagazine.MaritalStatus
			switch answers.MaritalStatus {
			case "Single":
				maritalStatus = dailylifemagazine.MARITALSTATUS_SINGLE
			case "Married":
				maritalStatus = dailylifemagazine.MARITALSTATUS_MARRIED
			case "Widowed":
				maritalStatus = dailylifemagazine.MARITALSTATUS_WIDOWED
			case "Divorced":
				maritalStatus = dailylifemagazine.MARITALSTATUS_DIVORCED
			default:
				panic(errInvalidMaritalStatus)
			}

			readers = append(readers, dailylifemagazine.NewReader(answers.Age, gender, maritalStatus, answers.AnnualIncome))
			if errAskAddAnotherReader := survey.AskOne(
				&survey.Confirm{
					Message: "Would you like to add another reader?",
					Default: true,
				},
				&addingReaders,
			); errAskAddAnotherReader != nil {
				panic(errAskAddAnotherReader)
			}
		}

		if errSave := saveRecords(readers); errSave != nil {
			panic(errSave)
		}
	case "View readers by age group...":
		readers, errLoad := loadRecords()
		if errLoad != nil {
			panic(errLoad)
		} else if len(readers) == 0 {
			fmt.Println("No records")
			os.Exit(0)
		}

		var countUnder20, count20, count30, count40, count50Over int
		for _, reader := range readers {
			switch age := reader.Age(); {
			case age < 20:
				countUnder20++
			case age >= 20 && age < 30:
				count20++
			case age >= 30 && age < 40:
				count30++
			case age >= 40 && age < 50:
				count40++
			case age >= 50:
				count50Over++
			}
		}

		fmt.Printf("Under 20:      %d\n", countUnder20)
		fmt.Printf("Between 20-29: %d\n", count20)
		fmt.Printf("Between 30-39: %d\n", count30)
		fmt.Printf("Between 40-49: %d\n", count40)
		fmt.Printf("50 and Over:   %d\n", count50Over)
	case "View readers by gender and age group...":
		readers, errLoad := loadRecords()
		if errLoad != nil {
			panic(errLoad)
		} else if len(readers) == 0 {
			fmt.Println("No records")
			os.Exit(0)
		}

		var maleCountUnder20, maleCount20, maleCount30, maleCount40, maleCount50Over int
		var femaleCountUnder20, femaleCount20, femaleCount30, femaleCount40, femaleCount50Over int
		for _, reader := range readers {
			switch reader.Gender() {
			case dailylifemagazine.GENDER_MALE:
				switch age := reader.Age(); {
				case age < 20:
					maleCountUnder20++
				case age >= 20 && age < 30:
					maleCount20++
				case age >= 30 && age < 40:
					maleCount30++
				case age >= 40 && age < 50:
					maleCount40++
				case age >= 50:
					maleCount50Over++
				}
			case dailylifemagazine.GENDER_FEMALE:
				switch age := reader.Age(); {
				case age < 20:
					femaleCountUnder20++
				case age >= 20 && age < 30:
					femaleCount20++
				case age >= 30 && age < 40:
					femaleCount30++
				case age >= 40 && age < 50:
					femaleCount40++
				case age >= 50:
					femaleCount50Over++
				}
			}
		}

		fmt.Println("Under 20:")
		fmt.Printf("\tMale:   %d\n", maleCountUnder20)
		fmt.Printf("\tFemale: %d\n", femaleCountUnder20)
		fmt.Println("Between 20 - 29:")
		fmt.Printf("\tMale:   %d\n", maleCount20)
		fmt.Printf("\tFemale: %d\n", femaleCount20)
		fmt.Println("Between 30 - 39:")
		fmt.Printf("\tMale:   %d\n", maleCount30)
		fmt.Printf("\tFemale: %d\n", femaleCount30)
		fmt.Println("Between 40 - 49:")
		fmt.Printf("\tMale:   %d\n", maleCount40)
		fmt.Printf("\tFemale: %d\n", femaleCount40)
		fmt.Println("50 and Over:")
		fmt.Printf("\tMale:   %d\n", maleCount50Over)
		fmt.Printf("\tFemale: %d\n", femaleCount50Over)
	case "View readers by annual income...":
		readers, errLoad := loadRecords()
		if errLoad != nil {
			panic(errLoad)
		} else if len(readers) == 0 {
			fmt.Println("No records")
			os.Exit(0)
		}

		var countUnder30K, count30Kto50K, count50Kto70K, countOver70K int
		for _, reader := range readers {
			switch annualIncome := reader.AnnualIncome(); {
			case annualIncome < 30000:
				countUnder30K++
			case annualIncome >= 30000 && annualIncome < 50000:
				count30Kto50K++
			case annualIncome >= 50000 && annualIncome < 70000:
				count50Kto70K++
			case annualIncome >= 70000:
				countOver70K++
			}
		}

		fmt.Printf("Under $30,000:             %d\n", countUnder30K)
		fmt.Printf("Between $30,000 - $49,000: %d\n", count30Kto50K)
		fmt.Printf("Between $50,000 - $69,000: %d\n", count50Kto70K)
		fmt.Printf("$70,000 and Over:          %d\n", countOver70K)
	}
}

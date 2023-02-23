package main

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sorucoder/southhills-it100-finals/chapter4/choretail"
)

var (
	errInvalidPrompt      error = errors.New("invalid prompt")
	errInvalidName        error = errors.New("invalid name")
	errInvalidTransaction error = errors.New("invalid transaction")
)

var nameRegexp *regexp.Regexp = regexp.MustCompile(`^[A-Za-z, -]+$`)

func validateName(answer any) error {
	if answerValue, ok := answer.(string); ok {
		if nameRegexp.MatchString(answerValue) {
			return nil
		}
		return errInvalidName
	}
	return errInvalidPrompt
}

func transformName(answer any) any {
	if answerValue, ok := answer.(string); ok {
		if nameRegexp.MatchString(answerValue) {
			var capitalized strings.Builder
			names := strings.Split(answerValue, " ")
			for index, name := range names {
				capitalized.WriteString(strings.ToUpper(name[0:1]) + strings.ToLower(name[1:]))
				if index < len(names)-1 {
					capitalized.WriteRune(' ')
				}
			}
			return capitalized.String()
		}
		panic(errInvalidName)
	}
	panic(errInvalidPrompt)
}

func validateTransaction(answer any) error {
	if answerValue, ok := answer.(string); ok {
		if answerValue == "next" || answerValue == "done" {
			return nil
		} else if _, errParse := strconv.ParseFloat(answerValue, 64); errParse == nil {
			return nil
		}
		return errInvalidTransaction
	}
	return errInvalidPrompt
}

func main() {
	var menuOption string
	if errAskMenuQuestion := survey.AskOne(
		&survey.Select{
			Message: "Please select an option:",
			Options: []string{
				"Add an Employee Sales Record...",
				"View Productivity Scores...",
				"View Bonuses...",
			},
		},
		&menuOption,
	); errAskMenuQuestion != nil {
		panic(errAskMenuQuestion)
	}

	switch menuOption {
	case "Add an Employee Sales Record...":
		var employeeName struct {
			FirstName string `survey:"firstName"`
			LastName  string `survey:"lastName"`
		}
		if errAskEmployeeName := survey.Ask(
			[]*survey.Question{
				{
					Name: "firstName",
					Prompt: &survey.Input{
						Message: "Please enter the employee's first name:",
						Help:    "Valid names can contain upper- and lower-case letters, spaces, and commas. Names will automatically be capitalized.",
					},
					Validate:  validateName,
					Transform: transformName,
				},
				{
					Name: "lastName",
					Prompt: &survey.Input{
						Message: "Please enter the employee's last name:",
						Help:    "Valid names can contain upper- and lower-case letters, spaces, and commas. Names will automatically be capitalized.",
					},
					Validate:  validateName,
					Transform: transformName,
				},
			},
			&employeeName,
		); errAskEmployeeName != nil {
			panic(errAskEmployeeName)
		}
		record := choretail.NewEmployeeSalesRecord(employeeName.FirstName, employeeName.LastName)

		var shiftsDone bool
		for !shiftsDone {
			var shiftDone bool
			shifts := record.Shifts()
			transactions := make([]float64, 0)
			for !shiftDone {
				var answer string
				if errAskTransaction := survey.AskOne(
					&survey.Input{
						Message: fmt.Sprintf(`Please enter transaction #%d for shift #%d:`, len(transactions)+1, shifts+1),
						Help:    `You can enter the transaction amount, "next" for the next shift, or "done" to complete the record.`,
					},
					&answer,
					survey.WithValidator(validateTransaction),
				); errAskTransaction != nil {
					panic(errAskTransaction)
				}

				if answer == "next" || answer == "done" {
					shiftDone = true
					if answer == "done" {
						shiftsDone = true
					}
				} else if transaction, errParse := strconv.ParseFloat(answer, 64); errParse == nil {
					transactions = append(transactions, transaction)
				}
			}
			record.AddShift(transactions)
		}

		if errSave := record.Save(); errSave != nil {
			panic(errSave)
		}
	case "View Productivity Scores...":
		if errWalk := filepath.WalkDir("./records/", func(path string, dirEntry fs.DirEntry, errPath error) error {
			if errPath != nil {
				return filepath.SkipDir
			}

			if !dirEntry.IsDir() {
				record, errLoad := choretail.LoadEmployeeSalesRecord(path)
				if errLoad != nil {
					return fmt.Errorf(`failed to open record "%s": %w`, path, errLoad)
				}

				score := record.ProductivityScore()
				if score < 50 {
					fmt.Printf("%s\t%.2f\n", record.Name(), score)
				} else {
					fmt.Printf("%s\t%.2f***\n", record.Name(), score)
				}
			}

			return nil
		}); errWalk != nil {
			panic(errWalk)
		}
	case "View Bonuses...":
		if errWalk := filepath.WalkDir("records", func(path string, dirEntry fs.DirEntry, errPath error) error {
			if errPath != nil {
				return filepath.SkipDir
			}

			if !dirEntry.IsDir() {
				record, errLoad := choretail.LoadEmployeeSalesRecord(path)
				if errLoad != nil {
					return fmt.Errorf(`failed to open record "%s": %w`, path, errLoad)
				}
				fmt.Printf("%s\t%.2f\n", record.Name(), record.Bonus())
			}

			return nil
		}); errWalk != nil {
			panic(errWalk)
		}
	}
}

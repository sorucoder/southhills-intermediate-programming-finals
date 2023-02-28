package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

var (
	errInvalidPrompt error = errors.New("invalid prompt")
	errInvalidGrade  error = errors.New("quiz scores are between 0 and 100")
)

func validateGrade(answer any) error {
	if answerValue, ok := answer.(string); ok {
		grade, errParse := strconv.ParseInt(answerValue, 10, 0)
		if errParse != nil {
			return errParse
		} else if grade < 0 || grade > 100 {
			return errInvalidGrade
		}
		return nil
	}
	return errInvalidPrompt
}

func sum(values []int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	return sum
}

func mean(values []int) int {
	return sum(values) / len(values)
}

func median(values []int) int {
	count := len(values)
	if count%2 == 0 {
		return (values[count/2-1] + values[count/2]) / 2
	} else {
		return values[count/2-1]
	}
}

func main() {
	var studentName string
	if errAskStudentName := survey.AskOne(
		&survey.Input{
			Message: "Please enter your name:",
		},
		&studentName,
	); errAskStudentName != nil {
		panic(errAskStudentName)
	}

	grades := make([]int, 0, 9)
	for quiz := 0; quiz < cap(grades); quiz++ {
		var gradeAnswer string
		if errAskQuizGrade := survey.AskOne(
			&survey.Input{
				Message: fmt.Sprintf(`Please enter the grade for quiz #%d:`, quiz+1),
			},
			&gradeAnswer,
			survey.WithValidator(survey.ComposeValidators(survey.Required, validateGrade)),
		); errAskQuizGrade != nil {
			panic(errAskQuizGrade)
		}
		grade, _ := strconv.ParseInt(gradeAnswer, 10, 0)
		grades = append(grades, int(grade))
	}
	sort.Ints(grades)
	grades = grades[3:]

	fmt.Println(studentName)
	fmt.Printf("Total Points:       %d\n", sum(grades))
	fmt.Printf("Mean Quiz Scores:   %d\n", mean(grades))
	fmt.Printf("Median Quiz Scores: %d\n", median(grades))
}

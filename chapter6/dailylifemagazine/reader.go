package dailylifemagazine

import (
	"encoding/json"
	"errors"
)

type Gender int

var ErrInvalidGender error = errors.New("invalid gender")

const (
	GENDER_MALE Gender = iota
	GENDER_FEMALE
)

func (gender Gender) String() string {
	switch gender {
	case GENDER_MALE:
		return "Male"
	case GENDER_FEMALE:
		return "Female"
	default:
		return "?"
	}
}

type MaritalStatus int

var ErrInvalidMaritalStatus error = errors.New("invalid marital status")

const (
	MARITALSTATUS_SINGLE MaritalStatus = iota
	MARITALSTATUS_MARRIED
	MARITALSTATUS_WIDOWED
	MARITALSTATUS_DIVORCED
)

func (maritalStatus MaritalStatus) String() string {
	switch maritalStatus {
	case MARITALSTATUS_SINGLE:
		return "Single"
	case MARITALSTATUS_MARRIED:
		return "Married"
	case MARITALSTATUS_WIDOWED:
		return "Widowed"
	case MARITALSTATUS_DIVORCED:
		return "Divorced"
	default:
		return "?"
	}
}

type readerJSON struct {
	Age           int           `json:"age"`
	Gender        Gender        `json:"gender"`
	MaritalStatus MaritalStatus `json:"marital_status"`
	AnnualIncome  int64         `json:"annual_income"`
}

type Reader struct {
	age           int
	gender        Gender
	maritalStatus MaritalStatus
	annualIncome  int64
}

func NewReader(age int, gender Gender, maritalStatus MaritalStatus, annualIncome int64) *Reader {
	return &Reader{
		age:           age,
		gender:        gender,
		maritalStatus: maritalStatus,
		annualIncome:  annualIncome,
	}
}

func (reader *Reader) Age() int {
	return reader.age
}

func (reader *Reader) Gender() Gender {
	return reader.gender
}

func (reader *Reader) MaritalStatus() MaritalStatus {
	return reader.maritalStatus
}

func (reader *Reader) AnnualIncome() int64 {
	return reader.annualIncome
}

func (reader *Reader) MarshalJSON() ([]byte, error) {
	var jsonData readerJSON
	jsonData.Age = reader.age
	jsonData.Gender = reader.gender
	jsonData.MaritalStatus = reader.maritalStatus
	jsonData.AnnualIncome = reader.annualIncome
	return json.Marshal(jsonData)
}

func (reader *Reader) UnmarshalJSON(jsonBytes []byte) error {
	var jsonData readerJSON
	if errUnmarshal := json.Unmarshal(jsonBytes, &jsonData); errUnmarshal != nil {
		return errUnmarshal
	}
	reader.age = jsonData.Age
	reader.gender = jsonData.Gender
	reader.maritalStatus = jsonData.MaritalStatus
	reader.annualIncome = jsonData.AnnualIncome
	return nil
}

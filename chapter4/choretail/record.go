package choretail

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type employeeSalesRecordJSON struct {
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Transactions [][]float64 `json:"transactions"`
}

type EmployeeSalesRecord struct {
	firstName    string
	lastName     string
	transactions [][]float64
}

func NewEmployeeSalesRecord(firstName, lastName string) *EmployeeSalesRecord {
	return &EmployeeSalesRecord{
		firstName:    firstName,
		lastName:     lastName,
		transactions: make([][]float64, 0),
	}
}

func LoadEmployeeSalesRecord(path string) (*EmployeeSalesRecord, error) {
	jsonFile, errOpen := os.Open(path)
	if errOpen != nil {
		return nil, errOpen
	}
	defer jsonFile.Close()

	var record EmployeeSalesRecord
	jsonDecoder := json.NewDecoder(jsonFile)
	if errDecode := jsonDecoder.Decode(&record); errDecode != nil {
		return nil, errDecode
	}

	return &record, nil
}

func (record *EmployeeSalesRecord) MarshalJSON() ([]byte, error) {
	var jsonData employeeSalesRecordJSON
	jsonData.FirstName = record.firstName
	jsonData.LastName = record.lastName
	jsonData.Transactions = record.transactions
	return json.Marshal(jsonData)
}

func (record *EmployeeSalesRecord) UnmarshalJSON(jsonBytes []byte) error {
	var jsonData employeeSalesRecordJSON
	if errUnmarshal := json.Unmarshal(jsonBytes, &jsonData); errUnmarshal != nil {
		return errUnmarshal
	}

	record.firstName = jsonData.FirstName
	record.lastName = jsonData.LastName
	record.transactions = jsonData.Transactions

	return nil
}

func (record *EmployeeSalesRecord) AddShift(transactions []float64) {
	record.transactions = append(record.transactions, transactions)
}

func (record *EmployeeSalesRecord) Name() string {
	return fmt.Sprintf(`%s %s`, record.firstName, record.lastName)
}

func (record *EmployeeSalesRecord) Shifts() int {
	return len(record.transactions)
}

func (record *EmployeeSalesRecord) Transactions() int {
	var transactions int
	for _, shift := range record.transactions {
		transactions += len(shift)
	}
	return transactions
}

func (record *EmployeeSalesRecord) Sales() float64 {
	var sales float64
	for _, shift := range record.transactions {
		for _, transaction := range shift {
			sales += transaction
		}
	}
	return sales
}

func (record *EmployeeSalesRecord) ProductivityScore() float64 {
	return record.Sales() / float64(record.Transactions()) / float64(record.Shifts())
}

func (record *EmployeeSalesRecord) Bonus() float64 {
	switch score := record.ProductivityScore(); {
	case score >= 200:
		return 200.00
	case score > 30 && score < 80:
		return 50.00
	case score >= 80 && score < 200:
		return 100.00
	default:
		return 25.00
	}
}

func (record *EmployeeSalesRecord) Save() error {
	jsonFile, errCreate := os.Create(filepath.Join("records", fmt.Sprintf(`%s %s.json`, record.lastName, record.firstName)))
	if errCreate != nil {
		return errCreate
	}
	defer jsonFile.Close()

	jsonEncoder := json.NewEncoder(jsonFile)
	if errEncode := jsonEncoder.Encode(record); errEncode != nil {
		return errEncode
	}

	return nil
}

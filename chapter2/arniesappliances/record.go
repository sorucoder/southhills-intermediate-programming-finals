package arniesappliances

import (
	"encoding/json"
)

type refrigeratorRecordJSON struct {
	Model        string  `json:"model"`
	CapacityFeet float64 `json:"capacity"`
}

type RefrigeratorRecord struct {
	model        string
	capacityFeet float64
}

func NewRefrigeratorRecord(model string, widthInches, heightInches, depthInches float64) *RefrigeratorRecord {
	return &RefrigeratorRecord{
		model:        model,
		capacityFeet: widthInches * heightInches * depthInches / 1728.0,
	}
}

func (record *RefrigeratorRecord) UnmarshalJSON(jsonBytes []byte) error {
	var jsonData refrigeratorRecordJSON
	if errUnmarshal := json.Unmarshal(jsonBytes, &jsonData); errUnmarshal != nil {
		return errUnmarshal
	}

	record.model = jsonData.Model
	record.capacityFeet = jsonData.CapacityFeet
	return nil
}

func (record *RefrigeratorRecord) MarshalJSON() ([]byte, error) {
	var jsonData refrigeratorRecordJSON
	jsonData.Model = record.model
	jsonData.CapacityFeet = record.capacityFeet
	return json.Marshal(jsonData)
}

func (record *RefrigeratorRecord) Model() string {
	return record.model
}

func (record *RefrigeratorRecord) CapacityFeet() float64 {
	return record.capacityFeet
}

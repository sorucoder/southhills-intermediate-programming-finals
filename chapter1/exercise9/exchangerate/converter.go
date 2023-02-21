package exchangerate

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type pairAPIResponse struct {
	baseUpdateAPIResponse
	BaseCode         string  `json:"base_code"`
	TargetCode       string  `json:"target_code"`
	ConversionRate   float64 `json:"conversion_rate"`
	ConversionResult float64 `json:"conversion_result"`
}

type Converter struct {
	baseCurrency, targetCurrency *Currency
	lastUpdate, nextUpdate       time.Time
	rate                         float64
}

func NewConverter(fromCurrency, toCurrency *Currency) (*Converter, error) {
	converter := &Converter{
		baseCurrency:   fromCurrency,
		targetCurrency: toCurrency,
	}

	if errUpdate := converter.Update(); errUpdate != nil {
		return nil, errUpdate
	}

	return converter, nil
}

func Convert(fromCurrency, toCurrency *Currency, amount float64) float64 {
	httpResponse, errGet := http.Get(apiURL.JoinPath("pair", fromCurrency.Code, toCurrency.Code, strconv.FormatFloat(amount, 'f', 4, 64)).String())
	if errGet != nil {
		panic(errGet)
	}

	var jsonData pairAPIResponse
	jsonDecoder := json.NewDecoder(httpResponse.Body)
	if errDecode := jsonDecoder.Decode(&jsonData); errDecode != nil {
		panic(errDecode)
	}

	return jsonData.ConversionResult
}

func (converter *Converter) BaseCurrency() *Currency {
	return converter.baseCurrency
}

func (converter *Converter) SetBaseCurrency(fromCurrency *Currency) error {
	converter.baseCurrency = fromCurrency
	return converter.ForceUpdate()
}

func (converter *Converter) TargetCurrency() *Currency {
	return converter.targetCurrency
}

func (converter *Converter) SetTargetCurrency(toCurrency *Currency) error {
	converter.targetCurrency = toCurrency
	return converter.ForceUpdate()
}

func (converter *Converter) LastUpdate() time.Time {
	return converter.lastUpdate
}

func (converter *Converter) NextUpdate() time.Time {
	return converter.nextUpdate
}

func (conveter *Converter) Rate() float64 {
	return conveter.rate
}

func (converter *Converter) ForceUpdate() error {
	httpResponse, errGet := http.Get(apiURL.JoinPath("pair", converter.baseCurrency.Code, converter.targetCurrency.Code).String())
	if errGet != nil {
		return errGet
	}

	var jsonData pairAPIResponse
	jsonDecoder := json.NewDecoder(httpResponse.Body)
	if errDecode := jsonDecoder.Decode(&jsonData); errDecode != nil {
		return errDecode
	}

	converter.lastUpdate = time.Unix(jsonData.TimeLastUpdateUnix, 0)
	converter.nextUpdate = time.Unix(jsonData.TimeNextUpdateUnix, 0)
	converter.rate = jsonData.ConversionRate
	return nil
}

func (converter *Converter) Update() error {
	if converter.nextUpdate.IsZero() || time.Now().After(converter.nextUpdate) {
		return converter.ForceUpdate()
	}
	return nil
}

func (converter *Converter) Convert(amount float64) float64 {
	if errUpdate := converter.Update(); errUpdate != nil {
		panic(errUpdate)
	}
	return converter.rate * amount
}

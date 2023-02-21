package exchangerate

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Currency struct {
	Code string
	Name string
}

func (currency *Currency) String() string {
	return fmt.Sprintf(`%s %s`, currency.Code, currency.Name)
}

var SupportedCurrencies []*Currency

func initSupportedCurrencies() error {
	httpResponse, errGet := http.Get(apiURL.JoinPath("codes").String())
	if errGet != nil {
		return errGet
	}

	var jsonData struct {
		baseAPIResponse
		SupportedCodes [][2]string `json:"supported_codes"`
	}
	jsonDecoder := json.NewDecoder(httpResponse.Body)
	if errDecode := jsonDecoder.Decode(&jsonData); errDecode != nil {
		return errDecode
	} else if errInvalidResponse := jsonData.Verify(); errInvalidResponse != nil {
		return errInvalidResponse
	}

	SupportedCurrencies = make([]*Currency, 0, len(jsonData.SupportedCodes))
	for _, currencyData := range jsonData.SupportedCodes {
		SupportedCurrencies = append(SupportedCurrencies, &Currency{Code: currencyData[0], Name: currencyData[1]})
	}

	return nil
}

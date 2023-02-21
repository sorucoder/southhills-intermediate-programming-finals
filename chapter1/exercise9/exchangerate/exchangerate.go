package exchangerate

import (
	"errors"
	"net/url"
	"os"
)

var apiURL *url.URL

var (
	ErrUnknownResponse     error = errors.New("unknown exchangerate API response")
	ErrNoAPIKey            error = errors.New("exchangerate API key is not set")
	ErrInvalidAPIKey       error = errors.New("exchangerate API key is invalid")
	ErrUnverifiedAccount   error = errors.New("exchangerate account not verified")
	ErrQuotaReached        error = errors.New("exchangerate account quota reached")
	ErrMalformedRequest    error = errors.New("malformed request for exchangerate API")
	ErrUnsupportedCurrency error = errors.New("unsupported currency for exchangerate API")
)

type baseAPIResponse struct {
	Result        string `json:"result"`
	Documentation string `json:"documentation"`
	TermsOfUse    string `json:"terms_of_use"`
	ErrorType     string `json:"error-type"`
}

func (response *baseAPIResponse) Verify() error {
	switch response.Result {
	case "success":
		return nil
	case "error":
		switch response.ErrorType {
		case "invalid-key":
			return ErrInvalidAPIKey
		case "inactive-account":
			return ErrUnverifiedAccount
		case "quota-reached":
			return ErrQuotaReached
		case "unsupported-code":
			return ErrUnsupportedCurrency
		case "malformed-request":
			return ErrMalformedRequest
		}
	}
	return ErrUnknownResponse
}

type baseUpdateAPIResponse struct {
	baseAPIResponse
	TimeLastUpdateUnix int64  `json:"time_last_update_unix"`
	TimeLastUpdateUTC  string `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64  `json:"time_next_update_unix"`
	TimeNextUpdateUTC  string `json:"time_next_update_utc"`
}

func initURL() error {
	var errParse error
	apiURL, errParse = url.Parse(`https://v6.exchangerate-api.com/v6/`)
	if errParse != nil {
		return errParse
	}
	apiKey, apiKeySet := os.LookupEnv("EXCHANGERATE_API_KEY")
	if !apiKeySet {
		return ErrNoAPIKey
	}
	apiURL = apiURL.JoinPath(apiKey)
	return nil
}

func init() {
	if errURL := initURL(); errURL != nil {
		panic(errURL)
	}

	if errSupportedCurrencies := initSupportedCurrencies(); errSupportedCurrencies != nil {
		panic(errSupportedCurrencies)
	}
}

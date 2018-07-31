package http_test

import (
	"encoding/json"
	"io/ioutil"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopee/exchangerate"
	handler "github.com/shopee/exchangerate/http"
	"github.com/shopee/exchangerate/mock"
	"github.com/shopee/testingutil"
)

func TestNewCurrencyExchangeHandler(t *testing.T) {
	es := &mock.ExchangeRateService{}
	exchangeRateHandler := handler.NewExchangeRateHandler(es)
	testingutil.Assert(t, exchangeRateHandler != nil, "ExchangeRate is nil")
}

func TestAddNewCurrencyExchange(t *testing.T) {
	i := int64(1)
	testCases := []struct {
		number                int
		as                    handler.ExchangeRateServicer
		requestPath           string
		expectedRequest       string
		expectedStatusCode    int
		expectedExchangeRates *exchangerate.ExchangeRateResponse
	}{
		{
			number:                1,
			as:                    &mock.ExchangeRateService{},
			requestPath:           "/api/v1/exchange_rate",
			expectedRequest:       `{"id_currency_exchange": 1, "rate": 9000}`,
			expectedStatusCode:    http.StatusOK,
			expectedExchangeRates: &exchangerate.ExchangeRateResponse{&i, 1, "SGD", "IDR", 9000, 0},
		},
		{
			number:                2,
			as:                    &mock.ExchangeRateServiceWithError{},
			requestPath:           "/api/v1/exchange_rate",
			expectedRequest:       `{"id_currency_exchange": 99, "rate": 9000}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedExchangeRates: nil,
		},
		{
			number:                3,
			as:                    &mock.ExchangeRateServiceWithError{},
			requestPath:           "/api/v1/exchange_rate",
			expectedRequest:       `{"id_currency_exchange": 99}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedExchangeRates: nil,
		},
	}

	for _, tc := range testCases {
		currencyExchangeHandler := handler.NewExchangeRateHandler(tc.as)

		router := gin.New()
		router.POST("/api/v1/exchange_rate", currencyExchangeHandler.AddNewExchangeRate)

		ts := httptest.NewServer(router)
		defer ts.Close()

		b := strings.NewReader(tc.expectedRequest)

		res, err := http.Post(ts.URL+tc.requestPath, "application/json", b)
		testingutil.Ok(t, err)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		testingutil.Ok(t, err)

		testingutil.Equals(t, tc.expectedStatusCode, res.StatusCode)

		var expectedResp string
		if tc.expectedExchangeRates != nil {
			expectedRespByte, err := json.Marshal(map[string]interface{}{"data": tc.expectedExchangeRates})
			testingutil.Ok(t, err)
			expectedResp = string(expectedRespByte)
		} else {
			if tc.number == 2 {
				expectedResp = `{"message":"Bad Request"}`
			} else {
				expectedResp = `{"message":"Key: 'ExchangeRateRequest.Rate' Error:Field validation for 'Rate' failed on the 'required' tag"}`
			}
		}
		testingutil.Equals(t, expectedResp, string(body))
	}
}

func TestListExchangeRateCurrencyExchange(t *testing.T) {
	testCases := []struct {
		number                int
		as                    handler.ExchangeRateServicer
		requestPath           string
		expectedRequest       string
		expectedStatusCode    int
		expectedExchangeRates *[]exchangerate.ExchangeRateResponse
	}{
		{
			number:             1,
			as:                 &mock.ExchangeRateService{},
			requestPath:        "/api/v1/exchange_rate/list",
			expectedRequest:    `{"date": "2018-07-30"}`,
			expectedStatusCode: http.StatusOK,
			expectedExchangeRates: &[]exchangerate.ExchangeRateResponse{
				exchangerate.ExchangeRateResponse{nil, 1, "SGD", "IDR", 9000, 0},
				exchangerate.ExchangeRateResponse{nil, 2, "USD", "IDR", 13000, 13440.1},
				exchangerate.ExchangeRateResponse{nil, 3, "YEN", "IDR", 130, 0},
				exchangerate.ExchangeRateResponse{nil, 4, "MYR", "IDR", 3500, 3514.4}},
		},
		{
			number:                2,
			as:                    &mock.ExchangeRateServiceWithError{},
			requestPath:           "/api/v1/exchange_rate/list",
			expectedRequest:       `{"date": 1234}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedExchangeRates: nil,
		},
	}

	for _, tc := range testCases {
		currencyExchangeHandler := handler.NewExchangeRateHandler(tc.as)

		router := gin.New()
		router.POST("/api/v1/exchange_rate/list", currencyExchangeHandler.ListExchangeRate)

		ts := httptest.NewServer(router)
		defer ts.Close()

		b := strings.NewReader(tc.expectedRequest)

		res, err := http.Post(ts.URL+tc.requestPath, "application/json", b)
		testingutil.Ok(t, err)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		testingutil.Ok(t, err)

		testingutil.Equals(t, tc.expectedStatusCode, res.StatusCode)

		var expectedResp string
		if tc.expectedExchangeRates != nil {
			expectedRespByte, err := json.Marshal(map[string]interface{}{"data": tc.expectedExchangeRates})
			testingutil.Ok(t, err)
			expectedResp = string(expectedRespByte)
		} else {
			expectedResp = `{"message":"json: cannot unmarshal number into Go struct field ExchangeRateRequestList.date of type string"}`
		}
		testingutil.Equals(t, expectedResp, string(body))
	}
}

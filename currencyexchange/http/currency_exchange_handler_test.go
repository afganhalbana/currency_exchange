package http_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopee/currencyexchange"
	handler "github.com/shopee/currencyexchange/http"
	"github.com/shopee/currencyexchange/mock"
	"github.com/shopee/testingutil"
)

func TestNewCurrencyExchangeHandler(t *testing.T) {
	cs := &mock.CurrencyExchangeService{}
	currencyExchangeHandler := handler.NewCurrencyExchangeHandler(cs)
	testingutil.Assert(t, currencyExchangeHandler != nil, "CurrencyExchange is nil")
}

func TestAddNewCurrencyExchange(t *testing.T) {

	testCases := []struct {
		number                    int
		as                        handler.CurrencyExchangeServicer
		requestPath               string
		expectedRequest           string
		expectedStatusCode        int
		expectedCurrencyExchanges *currencyexchange.CurrencyExchangeResponse
	}{
		{
			number:                    1,
			as:                        &mock.CurrencyExchangeService{},
			requestPath:               "/api/v1/currency_exchange",
			expectedRequest:           `{"from":"SGD", "to":"IDR"}`,
			expectedStatusCode:        http.StatusOK,
			expectedCurrencyExchanges: &currencyexchange.CurrencyExchangeResponse{1, "SGD", "IDR"},
		},
		{
			number:                    2,
			as:                        &mock.CurrencyExchangeServiceWithError{},
			requestPath:               "/api/v1/currency_exchange",
			expectedRequest:           `{"from":"SGD"}`,
			expectedStatusCode:        http.StatusBadRequest,
			expectedCurrencyExchanges: nil,
		},
		{
			number:                    3,
			as:                        &mock.CurrencyExchangeServiceWithError{},
			requestPath:               "/api/v1/currency_exchange",
			expectedRequest:           `{"from":"SGD", "to":"SGD"}`,
			expectedStatusCode:        http.StatusBadRequest,
			expectedCurrencyExchanges: nil,
		},
	}

	for _, tc := range testCases {
		currencyExchangeHandler := handler.NewCurrencyExchangeHandler(tc.as)

		router := gin.New()
		router.POST("/api/v1/currency_exchange", currencyExchangeHandler.AddNewCurrencyExchange)

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
		if tc.expectedCurrencyExchanges != nil {
			expectedRespByte, err := json.Marshal(map[string]interface{}{"data": tc.expectedCurrencyExchanges})
			testingutil.Ok(t, err)
			expectedResp = string(expectedRespByte)
		} else {
			if tc.number == 2 {
				expectedResp = `{"message":"Key: 'CurrencyExchangeRequest.To' Error:Field validation for 'To' failed on the 'required' tag"}`
			} else {
				expectedResp = `{"message":"Cannot use same currency"}`
			}
		}
		testingutil.Equals(t, expectedResp, string(body))
	}

}

func TestRemoveExistingCurrencyExchange(t *testing.T) {
	testCases := []struct {
		number                    int
		as                        handler.CurrencyExchangeServicer
		requestPath               string
		expectedStatusCode        int
		expectedCurrencyExchanges *currencyexchange.CurrencyExchangeResponse
	}{
		{
			number:                    1,
			as:                        &mock.CurrencyExchangeService{},
			requestPath:               "/api/v1/currency_exchange/1",
			expectedStatusCode:        http.StatusOK,
			expectedCurrencyExchanges: &currencyexchange.CurrencyExchangeResponse{1, "SGD", "IDR"},
		},
		{
			number:                    2,
			as:                        &mock.CurrencyExchangeServiceWithError{},
			requestPath:               "/api/v1/currency_exchange/20",
			expectedStatusCode:        http.StatusBadRequest,
			expectedCurrencyExchanges: nil,
		},
		{
			number:                    3,
			as:                        &mock.CurrencyExchangeServiceWithError{},
			requestPath:               "/api/v1/currency_exchange/a",
			expectedStatusCode:        http.StatusBadRequest,
			expectedCurrencyExchanges: nil,
		},
	}

	for _, tc := range testCases {

		currencyExchangeHandler := handler.NewCurrencyExchangeHandler(tc.as)

		fmt.Println(currencyExchangeHandler.RemoveExistingCurrencyExchange)

		router := gin.New()
		router.DELETE("/api/v1/currency_exchange/:id", currencyExchangeHandler.RemoveExistingCurrencyExchange)

		ts := httptest.NewServer(router)
		defer ts.Close()

		client := &http.Client{}
		req, err := http.NewRequest("DELETE", ts.URL+tc.requestPath, nil)
		testingutil.Ok(t, err)

		res, err := client.Do(req)
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		testingutil.Ok(t, err)

		testingutil.Equals(t, tc.expectedStatusCode, res.StatusCode)

		var expectedResp string
		if tc.expectedCurrencyExchanges != nil {
			expectedRespByte, err := json.Marshal(map[string]interface{}{"data": tc.expectedCurrencyExchanges})
			testingutil.Ok(t, err)
			expectedResp = string(expectedRespByte)
		} else {
			if tc.number == 2 {
				expectedResp = `{"message":"Invalid ID"}`
			} else {
				expectedResp = `{"message":"strconv.Atoi: parsing \"a\": invalid syntax"}`
			}
		}
		testingutil.Equals(t, expectedResp, string(body))
	}

}

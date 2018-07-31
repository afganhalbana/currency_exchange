package currencyexchange_test

import (
	"errors"
	"testing"

	"github.com/shopee/currencyexchange"
	"github.com/shopee/currencyexchange/mock"
	"github.com/shopee/testingutil"
)

func TestNewCurrencyExchangeService(t *testing.T) {
	cr := &mock.CurrencyExchangeRepository{}

	currencyExchangeService := currencyexchange.NewCurrencyExchangeService(cr)
	testingutil.Assert(t, currencyExchangeService != nil, "CurrencyExchange is nil")
}

func TestAddNewCurrencyExchange(t *testing.T) {
	testCases := []struct {
		cr               currencyexchange.CurrencyExchangeRepository
		req              currencyexchange.CurrencyExchangeRequest
		expectedResponse *currencyexchange.CurrencyExchangeResponse
		expectedError    error
	}{
		{
			cr:               &mock.CurrencyExchangeRepository{},
			req:              currencyexchange.CurrencyExchangeRequest{"USD", "IDR"},
			expectedResponse: nil,
			expectedError:    errors.New("Currency Exchange already exist"),
		},
		{
			cr:  &mock.CurrencyExchangeRepositoryWithError{},
			req: currencyexchange.CurrencyExchangeRequest{"PGD", "IDR"},
			expectedResponse: &currencyexchange.CurrencyExchangeResponse{
				ID:   1,
				From: "PGD",
				To:   "IDR",
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		cs := currencyexchange.NewCurrencyExchangeService(tc.cr)

		exchangeRate, err := cs.AddNewCurrencyExchange(tc.req)

		testingutil.Equals(t, tc.expectedResponse, exchangeRate)
		testingutil.Equals(t, tc.expectedError, err)
	}

}

func TestRemoveCurrencyExchange(t *testing.T) {
	testCases := []struct {
		cr               currencyexchange.CurrencyExchangeRepository
		req              int64
		expectedResponse *currencyexchange.CurrencyExchangeResponse
		expectedError    error
	}{
		{
			cr:  &mock.CurrencyExchangeRepository{},
			req: 1,
			expectedResponse: &currencyexchange.CurrencyExchangeResponse{
				ID:   1,
				From: "USD",
				To:   "IDR",
			},
			expectedError: nil,
		},
		{
			cr:  &mock.CurrencyExchangeRepository{},
			req: 99,
			expectedResponse: &currencyexchange.CurrencyExchangeResponse{
				ID:   1,
				From: "USD",
				To:   "IDR",
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		cs := currencyexchange.NewCurrencyExchangeService(tc.cr)

		exchangeRate, err := cs.RemoveCurrencyExchange(tc.req)

		testingutil.Equals(t, tc.expectedResponse, exchangeRate)
		testingutil.Equals(t, tc.expectedError, err)
	}
}

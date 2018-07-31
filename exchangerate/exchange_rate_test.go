package exchangerate_test

import (
	"errors"
	"testing"

	"github.com/shopee/currencyexchange"
	mockCurrencyExchange "github.com/shopee/currencyexchange/mock"
	"github.com/shopee/exchangerate"
	mockExchangeRate "github.com/shopee/exchangerate/mock"
	"github.com/shopee/testingutil"
)

// func TestNewTryoutService(t *testing.T) {
// 	tr := &mock.TryoutRepository{}
// 	ts := tryout.NewTryoutService(tr)
// 	testingutil.Assert(t, ts != nil, "NewTryoutService returns nil")
// }

func TestNewExchangeRateService(t *testing.T) {
	er := &mockExchangeRate.ExchangeRateRepository{}
	cr := &mockCurrencyExchange.CurrencyExchangeRepository{}

	exchangeRateService := exchangerate.NewExchangeRateService(er, cr)
	testingutil.Assert(t, exchangeRateService != nil, "ExchangeRate is nil")
}

func TestAddNewExchangeRate(t *testing.T) {
	idExchangeRate := int64(1)
	testCases := []struct {
		er               exchangerate.ExchangeRateRepository
		cr               currencyexchange.CurrencyExchangeRepository
		req              exchangerate.ExchangeRateRequest
		expectedResponse *exchangerate.ExchangeRateResponse
		expectedError    error
	}{
		{
			er:  &mockExchangeRate.ExchangeRateRepository{},
			cr:  &mockCurrencyExchange.CurrencyExchangeRepository{},
			req: exchangerate.ExchangeRateRequest{1, 9000},
			expectedResponse: &exchangerate.ExchangeRateResponse{
				IDExchangeRate:     &idExchangeRate,
				IDCurrencyExchange: 1,
				Rate:               9000,
				From:               "SGD",
				To:                 "IDR",
			},
			expectedError: nil,
		},
		{
			er:               &mockExchangeRate.ExchangeRateRepositoryWithError{},
			cr:               &mockCurrencyExchange.CurrencyExchangeRepositoryWithError{},
			req:              exchangerate.ExchangeRateRequest{1, 9000},
			expectedResponse: nil,
			expectedError:    errors.New("Internal Server Error"),
		},
	}

	for _, tc := range testCases {
		as := exchangerate.NewExchangeRateService(tc.er, tc.cr)

		exchangeRate, err := as.AddNewExchangeRate(tc.req)

		testingutil.Equals(t, tc.expectedResponse, exchangeRate)
		testingutil.Equals(t, tc.expectedError, err)
	}

}

func TestFindListExchangeRate(t *testing.T) {
	from := "2018-07-30"
	testCases := []struct {
		er               exchangerate.ExchangeRateRepository
		cr               currencyexchange.CurrencyExchangeRepository
		req              string
		expectedResponse *[]exchangerate.ExchangeRateResponse
		expectedError    error
	}{
		{
			er:  &mockExchangeRate.ExchangeRateRepository{},
			cr:  &mockCurrencyExchange.CurrencyExchangeRepository{},
			req: from,
			expectedResponse: &[]exchangerate.ExchangeRateResponse{
				exchangerate.ExchangeRateResponse{nil, 1, "SGD", "IDR", 1500, 2500},
				exchangerate.ExchangeRateResponse{nil, 2, "USD", "IDR", 1500, 2500},
				exchangerate.ExchangeRateResponse{nil, 3, "YEN", "IDR", 1500, 2500},
				exchangerate.ExchangeRateResponse{nil, 4, "MYR", "IDR", 1500, 2500},
			},
			expectedError: nil,
		},
		{
			er:               &mockExchangeRate.ExchangeRateRepositoryWithError{},
			cr:               &mockCurrencyExchange.CurrencyExchangeRepositoryWithError{},
			req:              "",
			expectedResponse: nil,
			expectedError:    errors.New("Internal Server Error"),
		},
	}

	for _, tc := range testCases {
		as := exchangerate.NewExchangeRateService(tc.er, tc.cr)

		exchangeRate, err := as.FindListExchangeRate(tc.req)

		testingutil.Equals(t, tc.expectedResponse, exchangeRate)
		testingutil.Equals(t, tc.expectedError, err)
	}
}

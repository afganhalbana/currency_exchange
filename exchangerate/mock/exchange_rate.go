package mock

import (
	"errors"
	"time"

	"github.com/shopee/exchangerate"
	"github.com/shopee/models"
)

type ExchangeRateRepository struct{}

type ExchangeRateRepositoryWithError struct{}

type ExchangeRateService struct{}

type ExchangeRateServiceWithError struct{}

func (es *ExchangeRateService) AddNewExchangeRate(req exchangerate.ExchangeRateRequest) (*exchangerate.ExchangeRateResponse, error) {
	idExchangeRate := int64(1)
	return &exchangerate.ExchangeRateResponse{
		IDExchangeRate:     &idExchangeRate,
		IDCurrencyExchange: 1,
		Rate:               req.Rate,
		From:               "SGD",
		To:                 "IDR",
	}, nil
}

func (es *ExchangeRateServiceWithError) AddNewExchangeRate(req exchangerate.ExchangeRateRequest) (*exchangerate.ExchangeRateResponse, error) {
	return nil, errors.New("Bad Request")
}

func (es *ExchangeRateService) FindListExchangeRate(from string) (*[]exchangerate.ExchangeRateResponse, error) {

	return &[]exchangerate.ExchangeRateResponse{
		exchangerate.ExchangeRateResponse{nil, 1, "SGD", "IDR", 9000, 0},
		exchangerate.ExchangeRateResponse{nil, 2, "USD", "IDR", 13000, 13440.1},
		exchangerate.ExchangeRateResponse{nil, 3, "YEN", "IDR", 130, 0},
		exchangerate.ExchangeRateResponse{nil, 4, "MYR", "IDR", 3500, 3514.4},
	}, nil
}

func (es *ExchangeRateServiceWithError) FindListExchangeRate(from string) (*[]exchangerate.ExchangeRateResponse, error) {
	return nil, errors.New("Bad Request")
}

func (er *ExchangeRateRepository) FindRateAverageWeeklyExchangeRateByID(int64, time.Time) (*float64, error) {
	result := float64(2500)
	return &result, nil
}

func (er *ExchangeRateRepositoryWithError) FindRateAverageWeeklyExchangeRateByID(int64, time.Time) (*float64, error) {
	return nil, errors.New("Internal Server Error")
}

func (er *ExchangeRateRepository) FindRateExchangeRateByID(int64, time.Time) (*int64, error) {
	result := int64(1500)
	return &result, nil
}

func (er *ExchangeRateRepositoryWithError) FindRateExchangeRateByID(int64, time.Time) (*int64, error) {
	return nil, errors.New("Internal Server Error")
}

func (er *ExchangeRateRepository) FindExchangeRateByID(int64) (*models.ExchangeRate, error) {
	return &models.ExchangeRate{
		models.Base{1, time.Now(), time.Now(), nil},
		models.CurrencyExchange{models.Base{1, time.Now(), time.Now(), nil}, "SGD", "IDR"},
		1, 9000}, nil
}

func (er *ExchangeRateRepositoryWithError) FindExchangeRateByID(int64) (*models.ExchangeRate, error) {
	return nil, errors.New("Internal Server Error")
}

func (er *ExchangeRateRepository) SaveExchangeRate(models.ExchangeRate) (*models.ExchangeRate, error) {
	return &models.ExchangeRate{
		models.Base{1, time.Now(), time.Now(), nil},
		models.CurrencyExchange{models.Base{1, time.Now(), time.Now(), nil}, "SGD", "IDR"},
		1, 9000}, nil
}

func (er *ExchangeRateRepositoryWithError) SaveExchangeRate(models.ExchangeRate) (*models.ExchangeRate, error) {
	return nil, errors.New("Internal Server Error")
}

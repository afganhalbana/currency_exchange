package mock

import (
	"errors"
	"time"

	"github.com/shopee/currencyexchange"
	"github.com/shopee/models"
)

type CurrencyExchangeRepository struct{}

type CurrencyExchangeRepositoryWithError struct{}

type CurrencyExchangeService struct{}

type CurrencyExchangeServiceWithError struct{}

func (cs *CurrencyExchangeService) AddNewCurrencyExchange(req currencyexchange.CurrencyExchangeRequest) (*currencyexchange.CurrencyExchangeResponse, error) {
	return &currencyexchange.CurrencyExchangeResponse{
		ID:   int64(1),
		From: "SGD",
		To:   "IDR",
	}, nil
}

func (cs *CurrencyExchangeServiceWithError) AddNewCurrencyExchange(req currencyexchange.CurrencyExchangeRequest) (*currencyexchange.CurrencyExchangeResponse, error) {
	return nil, errors.New("Bad Request")
}

func (cs *CurrencyExchangeService) RemoveCurrencyExchange(id int64) (*currencyexchange.CurrencyExchangeResponse, error) {
	return &currencyexchange.CurrencyExchangeResponse{
		ID:   int64(1),
		From: "SGD",
		To:   "IDR",
	}, nil
}

func (cs *CurrencyExchangeServiceWithError) RemoveCurrencyExchange(id int64) (*currencyexchange.CurrencyExchangeResponse, error) {
	return nil, errors.New("Bad Request")
}

func (cr *CurrencyExchangeRepository) SaveCurrencyExchange(param models.CurrencyExchange) (*models.CurrencyExchange, error) {
	return &models.CurrencyExchange{models.Base{1, time.Now(), time.Now(), nil}, param.From, param.To}, nil
}

func (cr *CurrencyExchangeRepositoryWithError) SaveCurrencyExchange(param models.CurrencyExchange) (*models.CurrencyExchange, error) {
	return nil, errors.New("Internal Server Error")
}

func (cr *CurrencyExchangeRepository) FindCurrencyExchangeByID(id int64) (*models.CurrencyExchange, error) {
	return &models.CurrencyExchange{models.Base{1, time.Now(), time.Now(), nil}, "SGD", "IDR"}, nil
}

func (cr *CurrencyExchangeRepositoryWithError) FindCurrencyExchangeByID(id int64) (*models.CurrencyExchange, error) {
	return nil, errors.New("Internal Server Error")
}

func (cr *CurrencyExchangeRepository) FindAllCurrencyExchange() (*[]models.CurrencyExchange, error) {
	return &[]models.CurrencyExchange{
		models.CurrencyExchange{models.Base{1, time.Now(), time.Now(), nil}, "SGD", "IDR"},
		models.CurrencyExchange{models.Base{2, time.Now(), time.Now(), nil}, "USD", "IDR"},
		models.CurrencyExchange{models.Base{3, time.Now(), time.Now(), nil}, "YEN", "IDR"},
		models.CurrencyExchange{models.Base{4, time.Now(), time.Now(), nil}, "MYR", "IDR"},
	}, nil
}

func (cr *CurrencyExchangeRepositoryWithError) FindAllCurrencyExchange() (*[]models.CurrencyExchange, error) {
	return nil, errors.New("Internal Server Error")
}

func (cr *CurrencyExchangeRepository) FindCurrencyExchangeFromTo(models.CurrencyExchange) (*models.CurrencyExchange, error) {
	return &models.CurrencyExchange{models.Base{1, time.Now(), time.Now(), nil}, "SGD", "IDR"}, nil
}

func (cr *CurrencyExchangeRepositoryWithError) FindCurrencyExchangeFromTo(models.CurrencyExchange) (*models.CurrencyExchange, error) {
	return nil, errors.New("records not found")
}

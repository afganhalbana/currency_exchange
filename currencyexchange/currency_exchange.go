package currencyexchange

import (
	"errors"
	"strings"
	"time"

	"github.com/shopee/models"
)

type (
	CurrencyExchangeRequest struct {
		From string `json:"from" binding:"required"`
		To   string `json:"to" binding:"required"`
	}

	CurrencyExchangeResponse struct {
		ID   int64  `json:"id" binding:"required"`
		From string `json:"from" binding:"required"`
		To   string `json:"to" binding:"required"`
	}

	CurrencyExchangeService struct {
		currencyExchangeRepository CurrencyExchangeRepository
	}

	CurrencyExchangeRepository interface {
		SaveCurrencyExchange(models.CurrencyExchange) (*models.CurrencyExchange, error)
		FindCurrencyExchangeByID(int64) (*models.CurrencyExchange, error)
		FindCurrencyExchangeFromTo(models.CurrencyExchange) (*models.CurrencyExchange, error)
		FindAllCurrencyExchange() (*[]models.CurrencyExchange, error)
	}
)

func NewCurrencyExchangeService(cr CurrencyExchangeRepository) *CurrencyExchangeService {
	return &CurrencyExchangeService{
		currencyExchangeRepository: cr,
	}
}

func (cs *CurrencyExchangeService) AddNewCurrencyExchange(req CurrencyExchangeRequest) (*CurrencyExchangeResponse, error) {
	currencyExchangeModel := models.CurrencyExchange{
		From: strings.ToUpper(req.From),
		To:   strings.ToUpper(req.To),
	}

	if isExist, err := cs.currencyExchangeRepository.FindCurrencyExchangeFromTo(currencyExchangeModel); (err != nil && err.Error() != "record not found") || isExist != nil {
		return nil, errors.New("Currency Exchange already exist")
	} else if response, err := cs.currencyExchangeRepository.SaveCurrencyExchange(currencyExchangeModel); err != nil {
		return nil, err
	} else {
		return &CurrencyExchangeResponse{
			ID:   response.ID,
			From: response.From,
			To:   response.To,
		}, nil
	}
}

func (cs *CurrencyExchangeService) RemoveCurrencyExchange(id int64) (*CurrencyExchangeResponse, error) {

	if existCurrencyExchange, err := cs.currencyExchangeRepository.FindCurrencyExchangeByID(id); err != nil {
		return nil, err
	} else {
		now := time.Now()
		existCurrencyExchange.DeletedAt = &now
		if response, err := cs.currencyExchangeRepository.SaveCurrencyExchange(*existCurrencyExchange); err != nil {
			return nil, err
		} else {
			return &CurrencyExchangeResponse{
				ID: response.ID,
			}, nil
		}
	}

}

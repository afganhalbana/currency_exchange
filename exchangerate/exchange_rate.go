package exchangerate

import (
	"time"

	"github.com/shopee/currencyexchange"
	"github.com/shopee/models"
)

type (
	ExchangeRateRequest struct {
		IDCurrencyExchange int64 `json:"id_currency_exchange" binding:"required"`
		Rate               int64 `json:"rate" binding:"required"`
	}

	ExchangeRateRequestList struct {
		Date string `json:"date"`
	}

	ExchangeRateResponse struct {
		IDExchangeRate     *int64  `json:"id_exchange_rate"`
		IDCurrencyExchange int64   `json:"id_currency_exchange"`
		From               string  `json:"from"`
		To                 string  `json:"to"`
		Rate               int64   `json:"rate"`
		WeeklyRate         float64 `json:"weekly_rate"`
	}

	ExchangeRateService struct {
		exchangeRateRepository     ExchangeRateRepository
		currencyExchangeRepository currencyexchange.CurrencyExchangeRepository
	}

	ExchangeRateRepository interface {
		FindRateAverageWeeklyExchangeRateByID(int64, time.Time) (*float64, error)
		FindRateExchangeRateByID(int64, time.Time) (*int64, error)
		FindExchangeRateByID(int64) (*models.ExchangeRate, error)
		SaveExchangeRate(models.ExchangeRate) (*models.ExchangeRate, error)
	}
)

func NewExchangeRateService(er ExchangeRateRepository, cr currencyexchange.CurrencyExchangeRepository) *ExchangeRateService {
	return &ExchangeRateService{
		exchangeRateRepository:     er,
		currencyExchangeRepository: cr,
	}
}

func (es *ExchangeRateService) AddNewExchangeRate(req ExchangeRateRequest) (*ExchangeRateResponse, error) {

	exchangeRateModel := models.ExchangeRate{
		IDCurrencyExchange: req.IDCurrencyExchange,
		Rate:               req.Rate,
	}

	currencyExchange, err := es.currencyExchangeRepository.FindCurrencyExchangeByID(req.IDCurrencyExchange)

	if err != nil {
		return nil, err
	}

	response, err := es.exchangeRateRepository.SaveExchangeRate(exchangeRateModel)

	if err != nil {
		return nil, err
	}

	return &ExchangeRateResponse{
		IDExchangeRate:     &response.ID,
		IDCurrencyExchange: response.IDCurrencyExchange,
		Rate:               response.Rate,
		From:               currencyExchange.From,
		To:                 currencyExchange.To,
	}, nil
}

func (es *ExchangeRateService) FindListExchangeRate(from string) (*[]ExchangeRateResponse, error) {
	if existCurrenctExchange, err := es.currencyExchangeRepository.FindAllCurrencyExchange(); err != nil {
		return nil, err
	} else {

		now := time.Now()

		if from != "" {
			const RFC3339FullDate = "2006-01-02"
			if now, err = time.Parse(RFC3339FullDate, from); err != nil {
				return nil, err
			}
		}

		var result []ExchangeRateResponse

		for _, currencyExchangeData := range *existCurrenctExchange {
			exchangeRateResponse := ExchangeRateResponse{
				From:               currencyExchangeData.From,
				To:                 currencyExchangeData.To,
				IDCurrencyExchange: currencyExchangeData.ID,
			}

			if averageRate, err := es.exchangeRateRepository.FindRateAverageWeeklyExchangeRateByID(currencyExchangeData.ID, now); (err != nil && err.Error() == "record not found") || averageRate == nil {
				// Insufficient Data
				exchangeRateResponse.WeeklyRate = 0
			} else {
				exchangeRateResponse.WeeklyRate = *averageRate
			}

			if rate, err := es.exchangeRateRepository.FindRateExchangeRateByID(currencyExchangeData.ID, now); (err != nil && err.Error() == "record not found") || rate == nil {
				exchangeRateResponse.Rate = 0
			} else {
				exchangeRateResponse.Rate = *rate
			}
			result = append(result, exchangeRateResponse)
		}

		return &result, nil
	}

}

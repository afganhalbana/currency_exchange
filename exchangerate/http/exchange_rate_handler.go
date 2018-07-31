package http

import (
	"github.com/gin-gonic/gin"
	"github.com/shopee/exchangerate"
	"github.com/shopee/responseutil"
)

type (
	ExchangeRateHandler struct {
		exchageRateService ExchangeRateServicer
	}

	ExchangeRateServicer interface {
		AddNewExchangeRate(exchangerate.ExchangeRateRequest) (*exchangerate.ExchangeRateResponse, error)
		FindListExchangeRate(string) (*[]exchangerate.ExchangeRateResponse, error)
	}
)

func NewExchangeRateHandler(es ExchangeRateServicer) *ExchangeRateHandler {
	return &ExchangeRateHandler{es}
}

func (eh *ExchangeRateHandler) AddNewExchangeRate(c *gin.Context) {
	var request exchangerate.ExchangeRateRequest

	if err := c.BindJSON(&request); err != nil {
		responseutil.BadRequest(c, err.Error())
	} else if result, err := eh.exchageRateService.AddNewExchangeRate(request); err != nil {
		responseutil.BadRequest(c, err.Error())
	} else {
		responseutil.OK(c, result)
	}
}

func (eh *ExchangeRateHandler) ListExchangeRate(c *gin.Context) {
	var request exchangerate.ExchangeRateRequestList

	if err := c.BindJSON(&request); err != nil {
		responseutil.BadRequest(c, err.Error())
	} else if result, err := eh.exchageRateService.FindListExchangeRate(request.Date); err != nil {
		responseutil.BadRequest(c, err.Error())
	} else {
		responseutil.OK(c, result)
	}
}

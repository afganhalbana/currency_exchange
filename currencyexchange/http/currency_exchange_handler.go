package http

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopee/currencyexchange"
	"github.com/shopee/responseutil"
)

type (
	CurrencyExchangeHandler struct {
		currencyExchangeService CurrencyExchangeServicer
	}

	CurrencyExchangeServicer interface {
		AddNewCurrencyExchange(currencyexchange.CurrencyExchangeRequest) (*currencyexchange.CurrencyExchangeResponse, error)
		RemoveCurrencyExchange(int64) (*currencyexchange.CurrencyExchangeResponse, error)
	}
)

func NewCurrencyExchangeHandler(cs CurrencyExchangeServicer) *CurrencyExchangeHandler {
	return &CurrencyExchangeHandler{cs}
}

func (ch *CurrencyExchangeHandler) AddNewCurrencyExchange(c *gin.Context) {
	var params currencyexchange.CurrencyExchangeRequest

	if err := c.BindJSON(&params); err != nil {
		responseutil.BadRequest(c, err.Error())
	} else if params.From == params.To {
		responseutil.BadRequest(c, "Cannot use same currency")
	} else if result, err := ch.currencyExchangeService.AddNewCurrencyExchange(params); err != nil {
		responseutil.BadRequest(c, err.Error())
	} else {
		responseutil.OK(c, result)
	}
}

func (ch *CurrencyExchangeHandler) RemoveExistingCurrencyExchange(c *gin.Context) {

	q := c.Params.ByName("id")

	if params, err := strconv.Atoi(q); err != nil {
		log.Print(err.Error())
		responseutil.BadRequest(c, err.Error())
	} else if result, err := ch.currencyExchangeService.RemoveCurrencyExchange(int64(params)); err != nil {
		responseutil.BadRequest(c, "Invalid ID")
	} else {
		responseutil.OK(c, result)
	}

}

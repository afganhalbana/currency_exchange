package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ce "github.com/shopee/currencyexchange"
	ceHandler "github.com/shopee/currencyexchange/http"
	"github.com/shopee/database"
	er "github.com/shopee/exchangerate"
	erHandler "github.com/shopee/exchangerate/http"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	db, err := database.New(database.Options{
		DBHost: viper.GetString(`database.host`),
		DBPort: viper.GetString(`database.port`),
		DBUser: viper.GetString(`database.user`),
		DBPass: viper.GetString(`database.pass`),
		DBName: viper.GetString(`database.name`),
	})

	if err != nil {
		panic(err)
		log.Fatal(err)
	}

	router := gin.Default()
	api := router.Group("/api")

	currencyExchangeRepository := database.NewCurrencyExchangeRepository(db)
	currencyExchangeService := ce.NewCurrencyExchangeService(currencyExchangeRepository)
	currencyExchangeHandler := ceHandler.NewCurrencyExchangeHandler(currencyExchangeService)

	exchangeRateRepository := database.NewExchangeRateRepository(db)
	exchangeRateService := er.NewExchangeRateService(exchangeRateRepository, currencyExchangeRepository)
	exchangeRateHandler := erHandler.NewExchangeRateHandler(exchangeRateService)

	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.POST("/v1/exchange_rate/", exchangeRateHandler.AddNewExchangeRate)
	api.POST("/v1/exchange_rate/list", exchangeRateHandler.ListExchangeRate)

	api.POST("/v1/currency_exchange/", currencyExchangeHandler.AddNewCurrencyExchange)
	api.DELETE("/v1/currency_exchange/:id", currencyExchangeHandler.RemoveExistingCurrencyExchange)

	router.Run(viper.GetString(`server.address`))

	defer db.Close()
}

func main() {

}

package models

type ExchangeRate struct {
	Base
	CurrencyExchange   CurrencyExchange
	IDCurrencyExchange int64 `gorm:"not null;type:bigint REFERENCES currency_exchanges(id)"`
	Rate               int64 `gorm:"not null;type:bigint"`
}

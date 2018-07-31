package models

type CurrencyExchange struct {
	Base
	From string `gorm:"not null;size:10"`
	To   string `gorm:"not null;size:10"`
}

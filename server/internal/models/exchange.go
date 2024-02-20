package models

/* Modelo auxiliar para extração dos dados da API de terceiro */
type Conversion struct {
	ExchangeRate ExchangeRate `json:"USDBRL"`
}

/* Modelo da taxa de conversão */
type ExchangeRate struct {
	Id  string `json:"-"`
	Bid string `json:"bid"`
}

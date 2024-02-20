package models

/*
	Modelo da taxa de conversão

Campo "error" é retornado apenas em casos de erro.
*/
type ExchangeRate struct {
	Bid   string `json:"bid,omitempty"`
	Error string `json:"error,omitempty"`
}

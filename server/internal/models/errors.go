package models

import "encoding/json"

/* Modelo auxiliar para gerar uma resposta padrão de erro */
type ErrorResponse struct {
	Error string `json:"error"`
}

/* Monta o JSON para resposta padrão de erro */
func ErrorResponseJson(err error) []byte {
	response := ErrorResponse{Error: err.Error()}
	data, _ := json.Marshal(response)
	return data
}

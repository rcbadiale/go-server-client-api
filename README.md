# Go Expert Server x Client API

Desafio de aplicação dos conhecimentos sobre webserver HTTP, contextos, banco
de dados e manipulação de arquivos em Go para a pós graduação Go Expert.

## Especificações

### Server

- Porta: 8080
- Endpoint: `GET /cotacao`
- Chama a API https://economia.awesomeapi.com.br/json/last/USD-BRL (timeout
200ms)
- Deve retornar o câmbio dólar x real (campo "bid")
- Salvar cotação num DB SQLite (timeout 10ms)
- Logs das falhas

Obs.: Foi implementada uma rota extra `GET /exchange/history` para facilitar a
visualização dos dados salvos no DB.

### Client

- Deverá chamar a rota `GET /cotacao` (timeout: 300ms)
- Salvar cotação no arquivo `cotacao.txt` (formato: `Dólar: {valor}`)
- Logs das falhas

## Execução

Necessário utilizar o Go 1.22.0.

### Server

Executar o comando: `go run server/cmd/main.go`

### Client

Executar o comando: `go run client/cmd/main.go`

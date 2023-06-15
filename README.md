# Criando a sua Primeira API Rest com Go

## DESCRIÇÃO DO DESAFIO:
Coloque seus conhecimentos de criação de servidores com GoLang e crie uma api rest utilizando todos os conceitos que você aprendeu até agora.

## Como executar:

API RESTful básica usando o roteador Gorilla Mux. Ele fornece os seguintes pontos finais:

- GET /parkingSpots: Obtém a lista de vagas de estacionamento disponíveis.
- POST /parkingSpots: Cria uma nova vaga de estacionamento.
- POST /cars: Estaciona um carro em uma vaga de estacionamento disponível.
- GET /cars/{id}: Obtém informações sobre um carro estacionado e libera a vaga de estacionamento.

- Você pode executar o código digitando no seu terminal:
```go
go run api.go
```
- API começará a ouvir em http://localhost:8080.
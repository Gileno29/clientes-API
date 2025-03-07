package handlers

import "github.com/Gileno29/clientes-API/models"

type ResponseCliente struct {
	Page     int              `json:"page"`
	Limit    int              `json:"limit"`
	Total    int64            `json:"total"`
	Clientes []models.Cliente `json:"clientes"`
}

type ResponseStatus struct {
	Uptime   float64
	Requests int
}

type ResponseErro struct {
	Mensagem string `json:"mensagem"` // Mensagem de erro
}

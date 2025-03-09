package repository

import (
	"github.com/Gileno29/clientes-API/dtos"
	"github.com/Gileno29/clientes-API/models"
)

type ClienteRepository interface {
	Create(cliente *models.Cliente) error
	FindByDocumento(documento string) (*models.Cliente, error)
	UpdateByDocumento(documento string, dadosAtualizados *dtos.AtualizaClienteRequest) (*models.Cliente, error)
	DeleteByDocumento(documento string) error
	ListarClientes(razaoSocial string, page, limit int) ([]models.Cliente, int64, error)
}

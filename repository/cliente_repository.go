package repository

import (
	"github.com/Gileno29/clientes-API/models"
)

type ClienteRepository interface {
	Create(cliente *models.Cliente) error
	FindByDocumento(documento string) (*models.Cliente, error)
	UpdateByDocumento(documento string, cliente *models.Cliente) error
	DeleteByDocumento(documento string) error
}

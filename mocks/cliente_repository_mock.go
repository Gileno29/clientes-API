package mocks

import (
	"github.com/Gileno29/clientes-API/dtos"
	"github.com/Gileno29/clientes-API/models"
	"github.com/stretchr/testify/mock"
)

type MockClienteRepository struct {
	mock.Mock
}

func (m *MockClienteRepository) UpdateByDocumento(documento string, dadosAtualizados *dtos.AtualizaClienteRequest) (*models.Cliente, error) {
	args := m.Called(documento, dadosAtualizados)
	return args.Get(0).(*models.Cliente), args.Error(1)
}

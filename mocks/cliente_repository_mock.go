package mocks

import (
	"github.com/Gileno29/clientes-API/dtos"
	"github.com/Gileno29/clientes-API/models"
	"github.com/stretchr/testify/mock"
)

type MockClienteRepository struct {
	mock.Mock
}

func (m *MockClienteRepository) Create(cliente *models.Cliente) error {
	args := m.Called(cliente)
	return args.Error(0)
}

func (m *MockClienteRepository) FindByDocumento(documento string) (*models.Cliente, error) {
	args := m.Called(documento)
	return args.Get(0).(*models.Cliente), args.Error(1)
}

func (m *MockClienteRepository) UpdateByDocumento(documento string, dadosAtualizados *dtos.AtualizaClienteRequest) (*models.Cliente, error) {
	args := m.Called(documento, dadosAtualizados)
	return args.Get(0).(*models.Cliente), args.Error(1)
}

func (m *MockClienteRepository) DeleteByDocumento(documento string) error {
	args := m.Called(documento)
	return args.Error(0)
}

func (m *MockClienteRepository) ListarClientes(razaoSocial string, page, limit int) ([]models.Cliente, int64, error) {
	args := m.Called(razaoSocial, page, limit)
	return args.Get(0).([]models.Cliente), args.Get(1).(int64), args.Error(2)
}

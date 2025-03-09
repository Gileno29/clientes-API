// repository/cliente_repository_impl.go
package repository

import (
	"github.com/Gileno29/clientes-API/models"
	"gorm.io/gorm"
)

type clienteRepository struct {
	db *gorm.DB
}

func NewClienteRepository(db *gorm.DB) ClienteRepository {
	return &clienteRepository{db: db}
}

func (r *clienteRepository) Create(cliente *models.Cliente) error {
	return r.db.Create(cliente).Error
}

func (r *clienteRepository) FindByDocumento(documento string) (*models.Cliente, error) {
	var cliente models.Cliente
	err := r.db.Where("documento = ?", documento).First(&cliente).Error
	if err != nil {
		return nil, err
	}
	return &cliente, nil
}

// UpdateByDocumento atualiza um cliente pelo documento
func (r *clienteRepository) UpdateByDocumento(documento string, cliente *models.Cliente) error {
	// Busca o cliente existente
	var existingCliente models.Cliente
	if err := r.db.Where("documento = ?", documento).First(&existingCliente).Error; err != nil {
		return err // Retorna erro se o cliente não for encontrado
	}

	// Atualiza os campos do cliente existente
	if cliente.RazaoSocial != "" {
		existingCliente.RazaoSocial = cliente.RazaoSocial
	}
	if cliente.Blocklist != existingCliente.Blocklist {
		existingCliente.Blocklist = cliente.Blocklist
	}

	// Salva as alterações no banco de dados
	return r.db.Save(&existingCliente).Error
}

// DeleteByDocumento remove um cliente pelo documento
func (r *clienteRepository) DeleteByDocumento(documento string) error {
	return r.db.Where("documento = ?", documento).Delete(&models.Cliente{}).Error
}

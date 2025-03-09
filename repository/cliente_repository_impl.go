// repository/cliente_repository_impl.go
package repository

import (
	"github.com/Gileno29/clientes-API/dtos"
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

func (r *clienteRepository) UpdateByDocumento(cliente *models.Cliente, dadosAtualizados *dtos.AtualizaClienteRequest) (*models.Cliente, error) {
	// Busca o cliente pelo documento

	// Atualiza os campos do cliente
	if dadosAtualizados.RazaoSocial != nil && *dadosAtualizados.RazaoSocial != "" && *dadosAtualizados.RazaoSocial != " " {
		cliente.RazaoSocial = *dadosAtualizados.RazaoSocial
	}
	if dadosAtualizados.Blocklist != nil {
		cliente.Blocklist = *dadosAtualizados.Blocklist
	}

	if err := r.db.Save(cliente).Error; err != nil {
		return nil, err
	}

	return cliente, nil
}

// DeleteByDocumento remove um cliente pelo documento
func (r *clienteRepository) DeleteByDocumento(documento string) error {
	return r.db.Where("documento = ?", documento).Delete(&models.Cliente{}).Error
}

func (r *clienteRepository) ListarClientes(razaoSocial string, page, limit int) ([]models.Cliente, int64, error) {
	var clientes []models.Cliente
	var total int64

	query := r.db.Model(&models.Cliente{})

	if razaoSocial != "" {
		query = query.Where("LOWER(razao_social) LIKE LOWER(?)", "%"+razaoSocial+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("razao_social ASC").Offset(offset).Limit(limit).Find(&clientes).Error; err != nil {
		return nil, 0, err
	}

	return clientes, total, nil
}

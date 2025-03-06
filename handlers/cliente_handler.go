package handlers

import (
	"net/http"

	"github.com/Gileno29/clientes-API/database"
	"github.com/Gileno29/clientes-API/models"
	"github.com/Gileno29/clientes-API/utils"

	"github.com/gin-gonic/gin"
)

func CadastrarCliente(c *gin.Context) {
	var cliente models.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Valida CPF/CNPJ
	if !utils.ValidarCPF(cliente.Documento) && !utils.ValidarCNPJ(cliente.Documento) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Documento inválido"})
		return
	}

	// Verifica se o cliente já existe
	var existingCliente models.Cliente
	if err := database.DB.Where("documento = ?", cliente.Documento).First(&existingCliente).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Cliente já cadastrado"})
		return
	}

	// Cria o cliente
	if err := database.DB.Create(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar cliente"})
		return
	}

	c.JSON(http.StatusCreated, cliente)
}

func ListarClientes(c *gin.Context) {
	var clientes []models.Cliente
	query := database.DB

	// Filtro por nome/razão social
	if nome := c.Query("nome"); nome != "" {
		query = query.Where("nome LIKE ?", "%"+nome+"%")
	}

	// Ordenação por nome
	query.Order("nome asc").Find(&clientes)

	c.JSON(http.StatusOK, clientes)
}

func VerificarCliente(c *gin.Context) {
	documento := c.Param("documento")

	var cliente models.Cliente
	if err := database.DB.Where("documento = ?", documento).First(&cliente).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"uptime":   "X segundos", // Implemente a lógica de up-time
		"requests": "X requisições", // Implemente a contagem de requisições
	})
}
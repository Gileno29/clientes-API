package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Gileno29/clientes-API/database"
	"github.com/Gileno29/clientes-API/middlewares"
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
	if !utils.ValidaDocumento(cliente.Documento) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Documento inválido"})
		return
	}

	// Verifica se o cliente já existe
	var existingCliente models.Cliente
	if err := database.DB.Where("documento = ?", cliente.Documento).First(&existingCliente).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Cliente já cadastrado"})
		return
	}

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
	if nome := c.Query("razao_social"); nome != "" {
		query = query.Where("razao_social LIKE ?", "%"+nome+"%")
	}

	// Ordenação por nome
	query.Order("razao_social asc").Find(&clientes)

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

func AtualizaCliente(c *gin.Context) {

	documento := c.Param("documento")
	documento = strings.ReplaceAll(documento, ".", "")
	documento = strings.ReplaceAll(documento, "-", "")
	documento = strings.ReplaceAll(documento, "/", "")
	documento = strings.TrimSpace(documento)

	fmt.Println("Esse é meu documento ", documento)
	var cliente models.Cliente
	if err := database.DB.Where("documento = ?", documento).First(&cliente).Error; err != nil {
		fmt.Println("Erro ao buscar cliente", err)
		c.JSON(http.StatusConflict, gin.H{"error": "Cliente Não identificado"})
		return
	}

	var dadosAtualizados struct {
		RazaoSocial *string `json:"razaosocial"`
		Blocklist   *bool   `json:"blocklist"`
	}

	if err := c.ShouldBindJSON(&dadosAtualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos, Algum parâmetro está vazio"})
		return
	}

	if dadosAtualizados.RazaoSocial != nil {
		cliente.RazaoSocial = *dadosAtualizados.RazaoSocial
	}
	if dadosAtualizados.Blocklist != nil {
		cliente.Blocklist = *dadosAtualizados.Blocklist
	}

	if err := database.DB.Save(&cliente).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente"})
		return
	}

	c.JSON(http.StatusOK, cliente)

}

func Status(c *gin.Context) {

	uptime := time.Since(utils.StartTime).Seconds()

	requests := middlewares.GetRequestCount()

	c.JSON(http.StatusOK, gin.H{
		"uptime":   uptime,
		"requests": requests,
	})
}

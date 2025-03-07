package handlers

import (
	"fmt"
	"net/http"
	"strconv"
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
		erro := ResponseErro{
			Mensagem: "{'error':" + err.Error() + "}",
		}
		c.JSON(http.StatusBadRequest, erro)
		return
	}

	// Valida CPF/CNPJ
	if !utils.ValidaDocumento(cliente.Documento) {
		erro := ResponseErro{
			Mensagem: "{'error': 'Documento inválido'}",
		}
		c.JSON(http.StatusBadRequest, erro)
		return
	}

	// Verifica se o cliente já existe
	var existingCliente models.Cliente
	if err := database.DB.Where("documento = ?", cliente.Documento).First(&existingCliente).Error; err == nil {
		erro := ResponseErro{
			Mensagem: "{'error': 'Cliente já cadastrado'}",
		}
		c.JSON(http.StatusConflict, erro)
		return
	}

	if err := database.DB.Create(&cliente).Error; err != nil {
		erro := ResponseErro{
			Mensagem: "{'error': 'Erro ao cadastrar cliente'}",
		}
		c.JSON(http.StatusInternalServerError, erro)
		return
	}

	c.JSON(http.StatusCreated, cliente)
}

// ListarClientes godoc
// @Summary Lista todos os clientes com paginação
// @Description Retorna uma lista de clientes com suporte a paginação e filtro por nome/razão social
// @Tags clientes
// @Accept json
// @Produce json
// @Param razao_social query string false "Filtrar por nome/razão social"
// @Param page query int false "Número da página" default(1)
// @Param limit query int false "Número de itens por página" default(10)
// @Success 200 {object} ResponseCliente "Resposta com clientes paginados"
// @Failure 400 {object} ResponseErro "Erro na requisição"
// @Failure 500 {object} ResponseErro "Erro interno do servidor"
// @Router /clientes [get]
func ListarClientes(c *gin.Context) {
	var clientes []models.Cliente
	query := database.DB

	if nome := c.Query("razao_social"); nome != "" {
		query = query.Where("razao_social LIKE ?", "%"+nome+"%")
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var total int64
	query.Model(&models.Cliente{}).Count(&total)

	query.Order("razao_social asc").Offset(offset).Limit(limit).Find(&clientes)

	resposta := ResponseCliente{
		Page:     page,
		Limit:    limit,
		Total:    total,
		Clientes: clientes,
	}

	c.JSON(http.StatusOK, resposta)
}

func VerificarCliente(c *gin.Context) {
	documento := c.Param("documento")

	if !utils.ValidaDocumento(documento) {
		erro := ResponseErro{
			Mensagem: "{'error': 'Documento inválido'}",
		}

		c.JSON(http.StatusBadRequest, erro)
		return
	}

	var cliente models.Cliente
	if err := database.DB.Where("documento = ?", documento).First(&cliente).Error; err != nil {
		erro := ResponseErro{
			Mensagem: "{'error': 'Cliente não encontrado'}",
		}
		c.JSON(http.StatusNotFound, erro)
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
		erro := ResponseErro{
			Mensagem: "{'error': 'Cliente Não identificado'}",
		}
		c.JSON(http.StatusConflict, erro)
		return
	}

	var dadosAtualizados struct {
		RazaoSocial *string `json:"razaosocial"`
		Blocklist   *bool   `json:"blocklist"`
	}

	if err := c.ShouldBindJSON(&dadosAtualizados); err != nil {
		erro := ResponseErro{
			Mensagem: "{'error': 'Dados inválidos, Algum parâmetro está vazio'}",
		}
		c.JSON(http.StatusBadRequest, erro)
		return
	}

	if dadosAtualizados.RazaoSocial != nil {
		cliente.RazaoSocial = *dadosAtualizados.RazaoSocial
	}
	if dadosAtualizados.Blocklist != nil {
		cliente.Blocklist = *dadosAtualizados.Blocklist
	}

	if err := database.DB.Save(&cliente).Error; err != nil {
		erro := ResponseErro{
			Mensagem: "{'error': 'Erro ao atualizar cliente'}",
		}
		c.JSON(http.StatusInternalServerError, erro)
		return
	}

	c.JSON(http.StatusOK, cliente)

}

func Status(c *gin.Context) {

	//uptime := time.Since(utils.StartTime).Seconds()

	//requests := middlewares.GetRequestCount()

	status := ResponseStatus{
		Uptime:   time.Since(utils.StartTime).Seconds(),
		Requests: middlewares.GetRequestCount(),
	}
	c.JSON(http.StatusOK, status)
}

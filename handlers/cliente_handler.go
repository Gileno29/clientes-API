package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Gileno29/clientes-API/dtos"
	"github.com/Gileno29/clientes-API/middlewares"
	"github.com/Gileno29/clientes-API/models"
	"github.com/Gileno29/clientes-API/repository"
	"github.com/Gileno29/clientes-API/utils"

	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	repo repository.ClienteRepository
}

func NewClienteHandler(repo repository.ClienteRepository) *ClienteHandler {
	return &ClienteHandler{repo: repo}
}

// @Summary Cadastra um novo cliente
// @Description Cadastra um novo cliente no sistema com base nos dados fornecidos.
// @Tags clientes
// @Accept json
// @Produce json
// @Param cliente body dtos.ClienteResponse true "Dados do cliente a ser cadastrado"
// @Success 201 {object} dtos.ClienteResponse "Cliente cadastrado com sucesso"
// @Failure 400 {object} dtos.ResponseErro "Erro ao processar a requisição (ex: documento inválido ou JSON inválido)"
// @Failure 409 {object} dtos.ResponseErro "Cliente já cadastrado"
// @Failure 500 {object} dtos.ResponseErro "Erro interno ao cadastrar o cliente"
// @Router /clientes [post]
func (h *ClienteHandler) CadastrarCliente(c *gin.Context) {
	var cliente models.Cliente

	if err := c.ShouldBindJSON(&cliente); err != nil {
		erro := dtos.ResponseErro{
			Mensagem: "{'error':" + err.Error() + "}",
		}
		c.JSON(http.StatusBadRequest, erro)
		return
	}

	// Valida CPF/CNPJ
	if !utils.ValidaDocumento(cliente.Documento) {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Documento inválido'}",
		}
		c.JSON(http.StatusBadRequest, erro)
		return
	}

	// Verifica se o cliente já existe
	existingCliente, err := h.repo.FindByDocumento(cliente.Documento)

	if err == nil && existingCliente != nil {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Cliente já cadastrado'}",
		}
		c.JSON(http.StatusConflict, erro)
		return
	}

	if err := h.repo.Create(&cliente); err != nil {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Erro ao cadastrar cliente'}",
		}
		c.JSON(http.StatusInternalServerError, erro)
		return
	}

	response := dtos.ClienteResponse{
		Documento:   cliente.Documento,
		RazaoSocial: cliente.RazaoSocial,
		Blocklist:   cliente.Blocklist,
	}

	c.JSON(http.StatusCreated, response)
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
// @Success 200 {object} dtos.ListarClientesResponse "Resposta com clientes paginados"
// @Failure 400 {object} dtos.ResponseErro "Erro na requisição"
// @Failure 500 {object} dtos.ResponseErro "Erro interno do servidor"
// @Router /clientes [get]
func (h *ClienteHandler) ListarClientes(c *gin.Context) {
	razaoSocial := c.Query("razao_social")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	clientes, total, err := h.repo.ListarClientes(razaoSocial, page, limit)

	if err != nil {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Erro ao listar clientes'}",
		}
		c.JSON(http.StatusInternalServerError, erro)
		return
	}

	if len(clientes) == 0 {
		erro := dtos.ResponseErro{
			Mensagem: "Nenhum cliente encontrado com o nome/razão social fornecido",
		}
		c.JSON(http.StatusNotFound, erro)
		return
	}

	var clientesResponse []dtos.ClienteResponse

	for _, cliente := range clientes {
		clientesResponse = append(clientesResponse, dtos.ClienteResponse{
			Documento:   cliente.Documento,
			RazaoSocial: cliente.RazaoSocial,
			Blocklist:   cliente.Blocklist,
		})
	}

	resposta := dtos.ListarClientesResponse{
		Page:     page,
		Limit:    limit,
		Total:    total,
		Clientes: clientesResponse,
	}

	c.JSON(http.StatusOK, resposta)
}

// VerificarCliente godoc
// @Summary Verifica se um cliente está cadastrado
// @Description Verifica se um cliente com o documento (CPF/CNPJ) fornecido está cadastrado na base de dados.
// @Tags clientes
// @Accept json
// @Produce json
// @Param documento path string true "Documento do cliente (CPF/CNPJ)"
// @Success 200 {object} dtos.ClienteResponse "Cliente encontrado"
// @Failure 400 {object} dtos.ResponseErro "Documento inválido"
// @Failure 404 {object} dtos.ResponseErro "Cliente não encontrado"
// @Router /clientes/{documento} [get]
func (h *ClienteHandler) VerificarCliente(c *gin.Context) {
	documento := utils.ClearNumber(c.Param("documento"))

	if !utils.ValidaDocumento(documento) {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Documento inválido'}",
		}

		c.JSON(http.StatusBadRequest, erro)
		return
	}

	cliente, err := h.repo.FindByDocumento(documento)
	if err != nil {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Cliente não encontrado'}",
		}
		c.JSON(http.StatusNotFound, erro)
		return
	}

	response := dtos.ClienteResponse{
		Documento:   cliente.Documento,
		RazaoSocial: cliente.RazaoSocial,
		Blocklist:   cliente.Blocklist,
	}

	c.JSON(http.StatusOK, response)
}

// AtualizaCliente godoc
// @Summary Atualiza os dados de um cliente
// @Description Atualiza a razão social e/ou o status de blocklist de um cliente com base no documento (CPF/CNPJ) fornecido.
// @Tags clientes
// @Accept json
// @Produce json
// @Param documento path string true "Documento do cliente (CPF/CNPJ)"
//
// @Param body body dtos.AtualizaClienteRequest true "Dados para atualização"
// @Success 200 {object} dtos.ClienteResponse "Cliente atualizado com sucesso"
// @Failure 400 {object} dtos.ResponseErro "Dados inválidos ou parâmetros vazios"
// @Failure 404 {object} dtos.ResponseErro "Cliente não encontrado"
// @Failure 409 {object} dtos.ResponseErro "Cliente não identificado"
// @Failure 500 {object} dtos.ResponseErro "Erro ao atualizar cliente"
// @Router /clientes/{documento} [put]
func (h *ClienteHandler) AtualizaCliente(c *gin.Context) {

	documento := utils.ClearNumber(c.Param("documento"))

	if !utils.ValidaDocumento(documento) {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Documento inválido'}",
		}

		c.JSON(http.StatusBadRequest, erro)
		return
	}

	var dadosAtualizados dtos.AtualizaClienteRequest
	if err := c.ShouldBindJSON(&dadosAtualizados); err != nil {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Dados inválidos, Algum parâmetro está vazio'}",
		}
		c.JSON(http.StatusBadRequest, erro)
		return
	}

	clienteAtualizado, err := h.repo.UpdateByDocumento(documento, &dadosAtualizados)
	if err != nil {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Erro ao atualizar cliente'}",
		}
		c.JSON(http.StatusInternalServerError, erro)
		return
	}

	response := dtos.ClienteResponse{
		Documento:   clienteAtualizado.Documento,
		RazaoSocial: clienteAtualizado.RazaoSocial,
		Blocklist:   clienteAtualizado.Blocklist,
	}

	c.JSON(http.StatusOK, response)

}

// DeletarCliente godoc
// @Summary Deleta um cliente
// @Description Deleta um cliente com base no documento (CPF/CNPJ) fornecido.
// @Tags clientes
// @Accept json
// @Produce json
// @Param documento path string true "Documento do cliente (CPF/CNPJ)"
// @Success 200 {object} dtos.ResponseSucesso "Cliente deletado com sucesso"
// @Failure 400 {object} dtos.ResponseErro "Documento inválido"
// @Failure 404 {object} dtos.ResponseErro "Cliente não encontrado"
// @Failure 500 {object} dtos.ResponseErro "Erro ao deletar cliente"
// @Router /clientes/{documento} [delete]
func (h *ClienteHandler) DeletarCliente(c *gin.Context) {
	documento := utils.ClearNumber(c.Param("documento"))

	if !utils.ValidaDocumento(documento) {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Documento inválido'}",
		}
		c.JSON(http.StatusBadRequest, erro)
		return
	}

	_, err := h.repo.FindByDocumento(documento)
	if err != nil {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Cliente não encontrado'}",
		}
		c.JSON(http.StatusNotFound, erro)
		return
	}

	if err := h.repo.DeleteByDocumento(documento); err != nil {
		erro := dtos.ResponseErro{
			Mensagem: "{'error': 'Erro ao deletar cliente'}",
		}
		c.JSON(http.StatusInternalServerError, erro)
		return
	}

	resposta := dtos.ResponseSucesso{
		Mensagem: "Cliente deletado com sucesso",
	}
	c.JSON(http.StatusOK, resposta)
}

// Status godoc
// @Summary Retorna o status do servidor
// @Description Retorna informações sobre o tempo de atividade (uptime) e o número de requisições atendidas.
// @Tags suporte
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseStatus "Status do servidor"
// @Router /status [get]
func Status(c *gin.Context) {

	//uptime := time.Since(utils.StartTime).Seconds()

	//requests := middlewares.GetRequestCount()

	status := dtos.ResponseStatus{
		Uptime:   time.Since(utils.StartTime).Seconds(),
		Requests: middlewares.GetRequestCount(),
	}
	c.JSON(http.StatusOK, status)
}

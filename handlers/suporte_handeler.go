package handlers

import (
	"net/http"
	"time"

	"github.com/Gileno29/clientes-API/dtos"
	"github.com/Gileno29/clientes-API/middlewares"
	"github.com/Gileno29/clientes-API/utils"
	"github.com/gin-gonic/gin"
)

// Status godoc
// @Summary Retorna o status do servidor
// @Description Retorna informações sobre o tempo de atividade (uptime) e o número de requisições atendidas.
// @Tags suporte
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseStatus "Status do servidor"
// @Router /status [get]

type SuporteHandler struct {
}

func NewSuporteHandler() *SuporteHandler {
	return &SuporteHandler{}
}

func (s *SuporteHandler) Status(c *gin.Context) {

	//uptime := time.Since(utils.StartTime).Seconds()

	//requests := middlewares.GetRequestCount()

	status := dtos.ResponseStatus{
		Uptime:   time.Since(utils.StartTime).Seconds(),
		Requests: middlewares.GetRequestCount(),
	}
	c.JSON(http.StatusOK, status)
}

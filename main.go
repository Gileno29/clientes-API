package main

import (
	"time"

	"github.com/Gileno29/clientes-API/database"
	_ "github.com/Gileno29/clientes-API/docs"
	"github.com/Gileno29/clientes-API/handlers"
	"github.com/Gileno29/clientes-API/middlewares"
	"github.com/Gileno29/clientes-API/repository"
	"github.com/Gileno29/clientes-API/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var startTime time.Time

func main() {

	utils.StartTime = time.Now()

	database.Connect()

	db := database.DB

	// Instancia repository e handler
	clienteRepo := repository.NewClienteRepository(db)
	clienteHandler := handlers.NewClienteHandler(clienteRepo)

	// instancia o GIN
	r := gin.Default()

	// configura ara utilizzar o midware para interceptação e contagem das requests.
	r.Use(middlewares.RequestCounterMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/clientes", clienteHandler.CadastrarCliente)
	r.GET("/clientes", clienteHandler.ListarClientes)
	r.GET("/clientes/:documento", clienteHandler.VerificarCliente)
	r.GET("/status", handlers.Status)
	r.PUT("/clientes/:documento", clienteHandler.AtualizaCliente)
	r.DELETE("/clientes/:documento", clienteHandler.DeletarCliente)
	r.Run(":8080")
}

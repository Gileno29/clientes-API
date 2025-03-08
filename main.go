package main

import (
	"time"

	"github.com/Gileno29/clientes-API/database"
	_ "github.com/Gileno29/clientes-API/docs"
	"github.com/Gileno29/clientes-API/handlers"
	"github.com/Gileno29/clientes-API/middlewares"
	"github.com/Gileno29/clientes-API/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var startTime time.Time

func main() {

	utils.StartTime = time.Now()

	database.Connect()

	r := gin.Default()

	r.Use(middlewares.RequestCounterMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/clientes", handlers.CadastrarCliente)
	r.GET("/clientes", handlers.ListarClientes)
	r.GET("/clientes/:documento", handlers.VerificarCliente)
	r.GET("/status", handlers.Status)
	r.PUT("/clientes/:documento", handlers.AtualizaCliente)
	r.DELETE("/clientes/:documento", handlers.DeletarCliente)
	r.Run(":8080")
}

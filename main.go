package main

import (
	"github.com/Gileno29/clientes-API/database"
	"github.com/Gileno29/clientes-API/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r := gin.Default()

	r.POST("/clientes", handlers.CadastrarCliente)
	r.GET("/clientes", handlers.ListarClientes)
	r.GET("/clientes/:documento", handlers.VerificarCliente)
	r.GET("/status", handlers.Status)
	r.PUT("/clientes/:documento", handlers.AtualizaCliente)
	r.Run(":8080")
}

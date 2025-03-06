package main

import (
	"cliente-api/database"
	"cliente-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r := gin.Default()


	r.POST("/clientes", handlers.CadastrarCliente)
	r.GET("/clientes", handlers.ListarClientes)
	r.GET("/clientes/:documento", handlers.VerificarCliente)
	r.GET("/status", handlers.Status)
	r.Run(":8080")
}
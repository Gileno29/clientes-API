package utils

import (
	"github.com/Gileno29/clientes-API/models" 
	"log"

	"gorm.io/gorm"
)

func VerificarTabelaClientes(db *gorm.DB) error {
	
	if !db.Migrator().HasTable(&models.Cliente{}) {
		log.Println("Tabela de clientes não encontrada. Criando tabela...")


		err := db.AutoMigrate(&models.Cliente{})
		if err != nil {
			log.Fatalf("Erro ao criar tabela de clientes: %v", err)
			return err
		}

		log.Println("Tabela de clientes criada com sucesso!")
	} else {
		log.Println("Tabela de clientes já existe.")
	}

	return nil
}
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Gileno29/clientes-API/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env: ", err)
	}

	conectioString := "user=" + os.Getenv("POSTGRES_USER") + " dbname=" + os.Getenv("POSTGRES_DB") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " host=" + os.Getenv("DATABASE_HOST") + " sslmode=disable"
	fmt.Println(conectioString)
	db, err := gorm.Open(postgres.Open(conectioString), &gorm.Config{})

	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}
	err = utils.VerificarTabelaClientes(db)
	if err != nil {
		panic("Falha ao criar tabela de clientes")
	}

	DB = db
}

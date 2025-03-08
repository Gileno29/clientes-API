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

	var (
		user     string
		password string
		dbname   string
		host     string
	)

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		log.Fatal("Variável de ambiente 'ENVIRONMENT' não definida")
	}

	switch env {
	case "development":
		user = os.Getenv("DEV_POSTGRES_USER")
		password = os.Getenv("DEV_POSTGRES_PASSWORD")
		dbname = os.Getenv("DEV_POSTGRES_DB")
		host = os.Getenv("DEV_DATABASE_HOST")
	case "test":
		user = os.Getenv("TEST_POSTGRES_USER")
		password = os.Getenv("TEST_POSTGRES_PASSWORD")
		dbname = os.Getenv("TEST_POSTGRES_DB")
		host = os.Getenv("TEST_DATABASE_HOST")
	case "production":
		user = os.Getenv("PROD_POSTGRES_USER")
		password = os.Getenv("PROD_POSTGRES_PASSWORD")
		dbname = os.Getenv("PROD_POSTGRES_DB")
		host = os.Getenv("PROD_DATABASE_HOST")
	default:
		log.Fatalf("Ambiente desconhecido: %s", env)
	}

	conectioString := "user=" + user + " dbname=" + dbname + " password=" + password + " host=" + host + " sslmode=disable"
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

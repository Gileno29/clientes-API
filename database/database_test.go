package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Chama a função Connect para estabelecer a conexão com o banco de dados
	Connect()

	// Verifica se a conexão foi estabelecida com sucesso
	assert.NotNil(t, DB, "A conexão com o banco de dados não deve ser nula")

	// Verifica se a tabela de clientes foi criada
	var tableExists bool
	DB.Raw("SELECT EXISTS (SELECT FROM pg_tables WHERE tablename = 'clientes')").Scan(&tableExists)
	assert.True(t, tableExists, "A tabela 'clientes' deve existir no banco de dados")
}
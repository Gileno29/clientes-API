package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type DatabaseIntegrationTestSuite struct {
	suite.Suite
	db *gorm.DB
}



func (suite *DatabaseIntegrationTestSuite) SetupSuite() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		suite.T().Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Conecta ao banco de dados
	Connect()
	suite.db = DB

	
	// Verifica se a conexão foi estabelecida
	assert.NotNil(suite.T(), suite.db, "A conexão com o banco de dados não deve ser nula")
}

// TestCreateTable verifica se a tabela de clientes foi criada
func (suite *DatabaseIntegrationTestSuite) TestCreateTable() {
	var tableExists bool
	suite.db.Raw("SELECT EXISTS (SELECT FROM pg_tables WHERE tablename = 'clientes')").Scan(&tableExists)
	assert.True(suite.T(), tableExists, "A tabela 'clientes' deve existir no banco de dados")
}

// TestInsertAndQueryClient testa a inserção e consulta de um cliente
func (suite *DatabaseIntegrationTestSuite) TestInsertAndQueryClient() {
	// Insere um cliente
	cliente := struct {
		CPF_CNPJ    string `gorm:"column:cpf_cnpj"`
		Nome        string `gorm:"column:nome"`
		RazaoSocial string `gorm:"column:razao_social"`
		Blocklist   bool   `gorm:"column:blocklist"`
	}{
		CPF_CNPJ:    "123.456.789-09",
		Nome:        "João Silva",
		RazaoSocial: "",
		Blocklist:   false,
	}

	result := suite.db.Table("clientes").Create(&cliente)
	assert.NoError(suite.T(), result.Error, "Erro ao inserir cliente")

	// Consulta o cliente inserido
	var clienteConsultado struct {
		CPF_CNPJ    string `gorm:"column:cpf_cnpj"`
		Nome        string `gorm:"column:nome"`
		RazaoSocial string `gorm:"column:razao_social"`
		Blocklist   bool   `gorm:"column:blocklist"`
	}
	suite.db.Table("clientes").Where("cpf_cnpj = ?", "123.456.789-09").First(&clienteConsultado)

	// Verifica se os dados estão corretos
	assert.Equal(suite.T(), "João Silva", clienteConsultado.Nome, "O nome do cliente deve ser 'João Silva'")
}

// TearDownSuite é executado uma vez após todos os testes
func (suite *DatabaseIntegrationTestSuite) TearDownSuite() {
	// Limpa a tabela de clientes após os testes
	suite.db.Exec("DELETE FROM clientes")
}

// Executa a suite de testes
func TestDatabaseIntegrationSuite(t *testing.T) {
	suite.Run(t, new(DatabaseIntegrationTestSuite))
}
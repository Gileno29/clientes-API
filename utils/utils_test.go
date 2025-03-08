package utils

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/Gileno29/clientes-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestValidaDocumento(t *testing.T) {
	// Testes para CPF
	assert.True(t, ValidaDocumento("529.982.247-25"), "CPF válido com formatação")
	assert.True(t, ValidaDocumento("52998224725"), "CPF válido sem formatação")
	assert.False(t, ValidaDocumento("529.982.247-26"), "CPF inválido")
	assert.False(t, ValidaDocumento("111.111.111-11"), "CPF com todos os dígitos iguais")
	assert.False(t, ValidaDocumento("123"), "CPF com menos de 11 dígitos")

	// Testes para CNPJ
	assert.True(t, ValidaDocumento("33.000.167/0001-01"), "CNPJ válido com formatação")
	assert.True(t, ValidaDocumento("33000167000101"), "CNPJ válido sem formatação")
	assert.False(t, ValidaDocumento("33.000.167/0001-02"), "CNPJ inválido")
	assert.False(t, ValidaDocumento("11.111.111/1111-11"), "CNPJ com todos os dígitos iguais")
	assert.False(t, ValidaDocumento("1234567890123"), "CNPJ com menos de 14 dígitos")
}


func TestValidarCPF(t *testing.T) {
	assert.True(t, ValidarCPF("52998224725"), "CPF válido")
	assert.False(t, ValidarCPF("52998224726"), "CPF inválido")
	assert.False(t, ValidarCPF("11111111111"), "CPF com todos os dígitos iguais")
	assert.False(t, ValidarCPF("123"), "CPF com menos de 11 dígitos")
}

func TestValidarCNPJ(t *testing.T) {
	assert.True(t, ValidarCNPJ("33000167000101"), "CNPJ válido")
	assert.False(t, ValidarCNPJ("33000167000102"), "CNPJ inválido")
	assert.False(t, ValidarCNPJ("11111111111111"), "CNPJ com todos os dígitos iguais")
	assert.False(t, ValidarCNPJ("1234567890123"), "CNPJ com menos de 14 dígitos")
}


func TestCalcularDigitoVerificador(t *testing.T) {
	assert.Equal(t, 5, calcularDigitoVerificador("529982247", 10), "Cálculo do primeiro dígito verificador do CPF")
	assert.Equal(t, 2, calcularDigitoVerificador("5299822472", 11), "Cálculo do segundo dígito verificador do CPF")
}


func TestCalcularDigitoVerificadorCNPJ(t *testing.T) {
	assert.Equal(t, 0, calcularDigitoVerificadorCNPJ("330001670001", []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}), "Cálculo do primeiro dígito verificador do CNPJ")
	assert.Equal(t, 1, calcularDigitoVerificadorCNPJ("3300016700010", []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}), "Cálculo do segundo dígito verificador do CNPJ")
}



func TestTodosDigitosIguais(t *testing.T) {
	assert.True(t, todosDigitosIguais("11111111111"), "Todos os dígitos iguais")
	assert.False(t, todosDigitosIguais("12345678901"), "Dígitos diferentes")
}


// setupDB inicializa um banco de dados SQLite em memória para testes
func setupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	return db
}

// TestVerificarTabelaClientes testa a função VerificarTabelaClientes
func TestVerificarTabelaClientes(t *testing.T) {
	// Configura o banco de dados em memória
	db := setupDB()

	// Testa a criação da tabela quando ela não existe
	t.Run("Cria tabela quando não existe", func(t *testing.T) {
		err := VerificarTabelaClientes(db)
		assert.NoError(t, err, "Erro ao verificar/criar tabela de clientes")

		// Verifica se a tabela foi criada
		hasTable := db.Migrator().HasTable(&models.Cliente{})
		assert.True(t, hasTable, "A tabela de clientes deve existir após a criação")
	})

	// Testa o comportamento quando a tabela já existe
	t.Run("Não recria tabela quando já existe", func(t *testing.T) {
		err := VerificarTabelaClientes(db)
		assert.NoError(t, err, "Erro ao verificar tabela de clientes")

		// Verifica se a tabela ainda existe
		hasTable := db.Migrator().HasTable(&models.Cliente{})
		assert.True(t, hasTable, "A tabela de clientes deve continuar existindo")
	})
}
package utils

import (
	"log"
	"testing"

	"github.com/Gileno29/clientes-API/models"
	"github.com/stretchr/testify/assert"
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
	// Teste para o primeiro dígito verificador do CPF
	// Para o CPF "123456789", o primeiro dígito verificador é 0
	primeiroDigito := calcularDigitoVerificador("123456789", 10)
	assert.Equal(t, 0, primeiroDigito, "Cálculo do primeiro dígito verificador do CPF")

	// Teste para o segundo dígito verificador do CPF
	// Para o CPF "1234567890", o segundo dígito verificador é 9
	segundoDigito := calcularDigitoVerificador("1234567890", 11)
	assert.Equal(t, 9, segundoDigito, "Cálculo do segundo dígito verificador do CPF")
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

func TestClearNumber(t *testing.T) {
	// Casos de teste
	tests := []struct {
		name     string // Nome do teste
		input    string // Entrada da função
		expected string // Saída esperada
	}{
		{
			name:     "String sem caracteres especiais",
			input:    "123456789",
			expected: "123456789",
		},
		{
			name:     "String com pontos",
			input:    "123.456.789",
			expected: "123456789",
		},
		{
			name:     "String com hífens",
			input:    "123-456-789",
			expected: "123456789",
		},
		{
			name:     "String com barras",
			input:    "123/456/789",
			expected: "123456789",
		},
		{
			name:     "String com pontos, hífens e barras",
			input:    "123.456-789/000",
			expected: "123456789000",
		},
		{
			name:     "String vazia",
			input:    "",
			expected: "",
		},
		{
			name:     "String com apenas caracteres especiais",
			input:    "./-",
			expected: "",
		},
	}

	// Executa os testes
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClearNumber(tt.input)
			if result != tt.expected {
				t.Errorf("ClearNumber(%s) = %s; esperado %s", tt.input, result, tt.expected)
			}
		})
	}
}

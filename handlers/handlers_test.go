package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Gileno29/clientes-API/database"
	"github.com/Gileno29/clientes-API/dtos"
	"github.com/Gileno29/clientes-API/mocks"

	"github.com/Gileno29/clientes-API/models"
	"github.com/Gileno29/clientes-API/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupDB inicializa um banco de dados SQLite em memória para testes
func setupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}
	// Cria a tabela de clientes
	db.AutoMigrate(&models.Cliente{})
	return db
}

// setupRouter inicializa o router do Gin com os handlers de teste
func setupRouter(db *gorm.DB) *gin.Engine {
	database.DB = db
	clienteRepo := repository.NewClienteRepository(db)
	clienteHandler := NewClienteHandler(clienteRepo)

	suporteHandler := NewSuporteHandler()

	router := gin.Default()
	router.POST("/clientes", clienteHandler.CadastrarCliente)
	router.GET("/clientes", clienteHandler.ListarClientes)
	router.GET("/clientes/:documento", clienteHandler.VerificarCliente)
	router.PUT("/clientes/:documento", clienteHandler.AtualizaCliente)
	router.DELETE("/clientes/:documento", clienteHandler.DeletarCliente)
	router.GET("/status", suporteHandler.Status)
	return router
}

func clearTable(db *gorm.DB) {
	db.Exec("DELETE FROM clientes") // Limpa a tabela de clientes
}
func TestCadastrarCliente(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	clearTable(db)

	// Caso de sucesso: Cliente válido
	t.Run("Cadastra cliente com sucesso", func(t *testing.T) {
		body := `{"documento": "52998224725", "razao_social": "João Silva", "blocklist": false}`
		req, _ := http.NewRequest("POST", "/clientes", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code, "Status code deve ser 201")
	})

	// Caso de erro: Documento inválido
	t.Run("Retorna erro para documento inválido", func(t *testing.T) {
		body := `{"documento": "123", "razao_social": "João Silva", "blocklist": false}`
		req, _ := http.NewRequest("POST", "/clientes", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code, "Status code deve ser 400")
	})

	// Caso de erro: Cliente já cadastrado
	t.Run("Retorna erro para cliente já cadastrado", func(t *testing.T) {
		// Insere um cliente no banco de dados
		db.Create(&models.Cliente{Documento: "52998224725", RazaoSocial: "João Silva", Blocklist: false})

		body := `{"documento": "52998224725", "razao_social": "João Silva", "blocklist": false}`
		req, _ := http.NewRequest("POST", "/clientes", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusConflict, resp.Code, "Status code deve ser 409")
	})
}

func TestListarClientes(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)
	clearTable(db)

	// Insere alguns clientes no banco de dados
	db.Create(&models.Cliente{Documento: "52998224725", RazaoSocial: "João Silva", Blocklist: false})
	db.Create(&models.Cliente{Documento: "33000167000101", RazaoSocial: "Empresa XYZ", Blocklist: true})

	// Caso de sucesso: Listar clientes sem filtro
	t.Run("Lista clientes sem filtro", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/clientes", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200")
	})

	// Caso de sucesso: Listar clientes com filtro por razão social
	t.Run("Lista clientes com filtro por razão social", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/clientes?razao_social=João", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200")
	})

	// Caso de erro: Nenhum cliente encontrado
	t.Run("Retorna erro quando nenhum cliente é encontrado", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/clientes?razao_social=inexistente", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code, "Status code deve ser 404")
	})
}

func TestVerificarCliente(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)
	clearTable(db)

	// Insere um cliente no banco de dados
	db.Create(&models.Cliente{Documento: "52998224725", RazaoSocial: "João Silva", Blocklist: false})

	// Caso de sucesso: Cliente encontrado
	t.Run("Retorna cliente encontrado", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/clientes/52998224725", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200")
	})

	// Caso de erro: Cliente não encontrado
	t.Run("Retorna erro quando cliente não é encontrado", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/clientes/12345678909", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code, "Status code deve ser 404")
	})

	// Caso de erro: Documento inválido
	t.Run("Retorna erro para documento inválido", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/clientes/123", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code, "Status code deve ser 400")
	})
}

func TestAtualizaCliente(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	// Insere um cliente no banco de dados
	db.Create(&models.Cliente{Documento: "52998224725", RazaoSocial: "João Silva", Blocklist: false})

	// Caso de sucesso: Atualiza razão social e blocklist
	t.Run("Atualiza cliente com sucesso", func(t *testing.T) {
		body := `{"razao_social": "João da Silva", "blocklist": true}`
		req, _ := http.NewRequest("PUT", "/clientes/52998224725", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200")
	})

	t.Run("Retorna erro quando cliente não é encontrado", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/clientes/12345678909", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code, "Status code deve ser 404")
	})

	// Caso de erro: Dados inválidos (JSON malformado ou campos vazios)
	t.Run("Retorna erro quando dados são inválidos", func(t *testing.T) {
		// Corpo da requisição com tipo incorreto para razao_social
		body := `{"razao_social":"gileno", "blocklist": Frue}` // razao_social é um número, mas deveria ser uma string
		req, _ := http.NewRequest("PUT", "/clientes/52998224725", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Verifica o status code
		assert.Equal(t, http.StatusBadRequest, resp.Code, "Status code deve ser 400")

		// Verifica a mensagem de erro
		var erroResponse dtos.ResponseErro
		err := json.Unmarshal(resp.Body.Bytes(), &erroResponse)
		assert.NoError(t, err, "Erro ao decodificar a resposta")
		assert.Contains(t, erroResponse.Mensagem, "Dados inválidos", "Mensagem de erro deve indicar dados inválidos")
	})

	// Caso de erro: Erro ao atualizar cliente no banco de dados
	t.Run("Retorna erro ao atualizar cliente no banco de dados", func(t *testing.T) {
		// Simula um erro no repository (ex: banco de dados indisponível)
		// Aqui você pode usar um mock do repository para forçar um erro
		mockRepo := new(mocks.MockClienteRepository)

		clienteExistente := &models.Cliente{
			Documento:   "52998224725",
			RazaoSocial: "João Silva",
			Blocklist:   false,
		}
		mockRepo.On("FindByDocumento", "52998224725").Return(clienteExistente, nil)

		mockRepo.On("UpdateByDocumento", mock.Anything, mock.Anything).Return(nil, errors.New("erro ao atualizar cliente"))

		handler := NewClienteHandler(mockRepo)

		router := gin.Default()
		router.PUT("/clientes/:documento", handler.AtualizaCliente)

		body := `{"razao_social": "João da Silva", "blocklist": true}`
		req, _ := http.NewRequest("PUT", "/clientes/52998224725", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code, "Status code deve ser 500")
	})

	// Caso de sucesso: Atualiza cliente com sucesso
	t.Run("Atualiza cliente com sucesso", func(t *testing.T) {
		// Cria o mock do repository
		mockRepo := new(mocks.MockClienteRepository)

		// Configura o comportamento esperado do mock para FindByDocumento
		clienteExistente := &models.Cliente{
			Documento:   "52998224725",
			RazaoSocial: "João Silva",
			Blocklist:   false,
		}
		mockRepo.On("FindByDocumento", "52998224725").Return(clienteExistente, nil)

		// Configura o comportamento esperado do mock para UpdateByDocumento
		clienteAtualizado := &models.Cliente{
			Documento:   "52998224725",
			RazaoSocial: "João da Silva", // Razão social atualizada
			Blocklist:   true,            // Blocklist atualizado
		}
		mockRepo.On("UpdateByDocumento", clienteExistente, mock.Anything).Return(clienteAtualizado, nil)

		// Cria o handler com o mock do repository
		handler := NewClienteHandler(mockRepo)

		// Configura o router
		router := gin.Default()
		router.PUT("/clientes/:documento", handler.AtualizaCliente)

		// Cria a requisição
		body := `{"razao_social": "João da Silva", "blocklist": true}`
		req, _ := http.NewRequest("PUT", "/clientes/52998224725", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Executa a requisição
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Verifica o resultado
		assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200")

		// Verifica o corpo da resposta
		var response dtos.ClienteResponse
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err, "Erro ao decodificar a resposta")
		assert.Equal(t, "52998224725", response.Documento, "Documento deve ser igual")
		assert.Equal(t, "João da Silva", response.RazaoSocial, "Razão social deve ser atualizada")
		assert.True(t, response.Blocklist, "Blocklist deve ser true")

		// Garante que o mock foi chamado conforme o esperado
		mockRepo.AssertExpectations(t)
	})

}

func TestDeletarCliente(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

	// Insere um cliente no banco de dados
	db.Create(&models.Cliente{Documento: "52998224725", RazaoSocial: "João Silva", Blocklist: false})

	// Caso de sucesso: Deleta cliente
	t.Run("Deleta cliente com sucesso", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/clientes/52998224725", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200")
	})

	// Caso de erro: Cliente não encontrado
	t.Run("Retorna erro quando cliente não é encontrado", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/clientes/12345678909", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code, "Status code deve ser 404")
	})

	// Caso de erro: Documento inválido
	t.Run("Retorna erro para documento inválido", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/clientes/123", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code, "Status code deve ser 400")
	})

}

func TestStatus(t *testing.T) {
	router := setupRouter(setupDB())

	// Caso de sucesso: Retorna status do servidor
	t.Run("Retorna status do servidor", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/status", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code, "Status code deve ser 200")
	})
}

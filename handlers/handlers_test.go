package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Gileno29/clientes-API/database"
	"github.com/Gileno29/clientes-API/dtos"
	"github.com/Gileno29/clientes-API/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	router := gin.Default()
	router.POST("/clientes", CadastrarCliente)
	router.GET("/clientes", ListarClientes)
	router.GET("/clientes/:documento", VerificarCliente)
	router.PUT("/clientes/:documento", AtualizaCliente)
	router.DELETE("/clientes/:documento", DeletarCliente)
	router.GET("/status", Status)
	return router
}

func TestCadastrarCliente(t *testing.T) {
	db := setupDB()
	router := setupRouter(db)

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
		req, _ := http.NewRequest("GET", "/clientes/12345678901", nil)
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

	// Caso de erro: Cliente não encontrado
	t.Run("Retorna erro quando cliente não é encontrado", func(t *testing.T) {
		body := `{"razao_social": "João da Silva", "blocklist": true}`
		req, _ := http.NewRequest("PUT", "/clientes/12345678901", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusConflict, resp.Code, "Status code deve ser 409")
	})

	// Caso de erro: Documento inválido
	t.Run("Retorna erro para documento inválido", func(t *testing.T) {
		body := `{"razao_social": "João da Silva", "blocklist": true}`
		req, _ := http.NewRequest("PUT", "/clientes/123", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code, "Status code deve ser 400")
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
		req, _ := http.NewRequest("DELETE", "/clientes/12345678901", nil)
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
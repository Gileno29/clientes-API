package middlewares

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	requestCount int
	mu           sync.Mutex
)

// RequestCounterMiddleware incrementa o contador de requisições a cada chamada
func RequestCounterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		requestCount++
		mu.Unlock()
		c.Next() // Passa para o próximo handler
	}
}

func GetRequestCount() int {
	mu.Lock()
	defer mu.Unlock()
	return requestCount
}

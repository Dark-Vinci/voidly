package middleware

import "github.com/gin-gonic/gin"

func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (m *Middleware) Default() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (m *Middleware) GenerateCredentials() string {
	return ""
}

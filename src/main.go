package main

import (
	"ThaiLy/configs"
	graphql_config "ThaiLy/graphql"
	"ThaiLy/middlewares"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func main() {
	// Set gin to release mode
	gin.SetMode(gin.ReleaseMode)
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Load database
	configs.InitDB()
	// Tạo gin
	r := gin.Default()

	// Tạo HTTP handler cho GraphQL

	r.POST("/graphql", middlewares.RequireAuth, func(c *gin.Context) {
		handler := handler.New(&handler.Config{
			Schema: graphql_config.Config(),
		})

		// Tạo context mới và truyền thông tin tài khoản vào
		ctx := context.WithValue(c.Request.Context(), "account", c.MustGet("account"))
		c.Request = c.Request.WithContext(ctx)

		handler.ServeHTTP(c.Writer, c.Request)
	})

	// Khởi tạo server với Gin
	log.Println("Server running at http://localhost:8080/graphql")
	r.Run(":8080")
}
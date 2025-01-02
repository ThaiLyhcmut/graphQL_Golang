package main

import (
	"ThaiLy/configs"
	graphql_config "ThaiLy/graphql"
	"ThaiLy/middlewares"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
)

func main() {
	// Set gin to release mode
	gin.SetMode(gin.ReleaseMode)
	// Load .env file
	godotenv.Load()

	fmt.Println("MySQL Name: ", os.Getenv("MYSQL_NAME"))
	// Load database connection
	configs.InitDB()

	// Create Gin router
	r := gin.Default()

	// CORS middleware setup
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Chỉ cho phép origin cụ thể
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// GraphQL route
	r.POST("/graphql", middlewares.RequireAuth, func(c *gin.Context) {
		handler := handler.New(&handler.Config{
			Schema: graphql_config.Config(), // GraphQL schema
		})

		// Checking if "account" exists in the context
		account, exists := c.Get("account")
		if exists {
			// Create new context with account if available
			ctx := context.WithValue(c.Request.Context(), "account", account)
			c.Request = c.Request.WithContext(ctx)
		}

		// Process the GraphQL request
		handler.ServeHTTP(c.Writer, c.Request)
	})

	// Start server on port 8080
	log.Println("Server running at http://localhost:8080/graphql")
	r.Run(":8080")
}

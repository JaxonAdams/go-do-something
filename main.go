package main

import (
	"context"
	"go-do-something/database"
	"go-do-something/routes"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load AWS Config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-1"))
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config: %v", err)
	}

	// Configure database connection
	svc := database.ConfigureDBClient(&cfg)

	// Configure HTTP router
	r := gin.Default()

	// Define routes
	v1 := r.Group("/api/v1")
	routes.RegisterTodoRoutes(v1, svc)

	r.Run()
}

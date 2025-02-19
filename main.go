package main

import (
	"context"
	"fmt"
	"go-do-something/database"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	// Load AWS Config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-1"))
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config: %v", err)
	}

	svc := database.ConfigureDBClient(&cfg)

	fmt.Println(svc)
}

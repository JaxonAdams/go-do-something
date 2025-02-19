package main

import (
	"context"
	"go-do-something/database"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	// Load AWS Config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-1"))
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config: %v", err)
	}

	svc := database.ConfigureDBClient(&cfg)

	// Test inserting a record to the database
	err = database.PutTodoItem(
		svc,
		"jaxon.adams@loanpro.io",
		strconv.FormatInt(time.Now().UnixMilli(), 10),
		"Test Item",
		"Test this TODO item",
		"TODO",
		"2025-02-28",
	)
	if err != nil {
		log.Fatalf("Failed to put TODO item: %v", err)
	}
}

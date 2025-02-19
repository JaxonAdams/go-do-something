package main

import (
	"context"
	"fmt"
	"go-do-something/database"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {
	// Load AWS Config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-1"))
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config: %v", err)
	}

	svc := database.ConfigureDBClient(&cfg)

	// Test inserting a record to the database
	// err = database.PutTodoItem(
	// 	svc,
	// 	"jaxontest@example.com",
	// 	strconv.FormatInt(time.Now().UnixMilli(), 10),
	// 	"OTHER USER ITEM",
	// 	"This is another user's item.",
	// 	"TODO",
	// 	"2025-02-28",
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to put TODO item: %v", err)
	// }

	// Test updating a to-do item's status
	err = database.UpdateTodoItemStatus(svc, "jaxontest@example.com", "1739949328475", "done")
	if err != nil {
		log.Fatalf("Failed to update to-do item status: %v", err)
	}

	// Test fetching a user's to-do list
	items, err := database.GetUserTodoList(svc, "jaxon.adams@loanpro.io")
	if err != nil {
		log.Fatalf("Failed to get to-do list: %v", err)
	}

	fmt.Println("\nUser To-Do List:")
	for _, item := range items {
		title := "N/A"
		description := "N/A"

		if titleAttr, ok := item["Title"].(*types.AttributeValueMemberS); ok {
			title = titleAttr.Value
		}

		if descAttr, ok := item["Description"].(*types.AttributeValueMemberS); ok {
			description = descAttr.Value
		}

		fmt.Printf("\n\tTitle: %s\n\tDescription: %s\n", title, description)
	}
}

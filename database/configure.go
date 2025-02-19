package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func ConfigureDBClient(cfg *aws.Config) *dynamodb.Client {

	// Create DynamoDB client
	svc := dynamodb.NewFromConfig(*cfg)

	// Create a table if it does not yet exist
	if err := createTable(svc); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return svc

}

func createTable(svc *dynamodb.Client) error {
	_, err := svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String("TodoItems"),
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("UserID"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("TodoID"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("Status"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("DueDate"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("UserID"), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String("ItemID"), KeyType: types.KeyTypeRange}, // Sort key
		},
		BillingMode: types.BillingModePayPerRequest,
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("StatusDueDateIndex"),
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String("Status"), KeyType: types.KeyTypeHash},
					{AttributeName: aws.String("DueDate"), KeyType: types.KeyTypeRange},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
				ProvisionedThroughput: nil,
			}
		},
	})

	if err != nil {
		if isTableExistsError(err) {
			fmt.Println("Table already exists, skipping creation.")
			return nil
		}
		return err
	}

	fmt.Println("Table created successfully!")
	return nil
}

func isTableExistsError(err error) bool {
	var tableExistsErr *types.ResourceInUseException
	return err != nil && errors.As(err, &tableExistsErr)
}

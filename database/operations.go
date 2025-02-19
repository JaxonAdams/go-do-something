package database

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func PutTodoItem(svc *dynamodb.Client, userID, todoID, title, description, status, dueDate string) error {
	_, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("TodoItems"),
		Item: map[string]types.AttributeValue{
			"UserID":      &types.AttributeValueMemberS{Value: userID},
			"TodoID":      &types.AttributeValueMemberS{Value: todoID},
			"Title":       &types.AttributeValueMemberS{Value: title},
			"Description": &types.AttributeValueMemberS{Value: description},
			"Status":      &types.AttributeValueMemberS{Value: status},
			"DueDate":     &types.AttributeValueMemberS{Value: dueDate},
			"CreatedAt":   &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
		},
	})
	if err != nil {
		return err
	}

	fmt.Println("Todo item inserted successfully!")
	return nil
}

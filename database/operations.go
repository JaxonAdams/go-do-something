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

func DeleteTodoItem(svc *dynamodb.Client, userID, todoID string) error {
	_, err := svc.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("TodoItems"),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: userID},
			"TodoID": &types.AttributeValueMemberS{Value: todoID},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserTodoList(svc *dynamodb.Client, userID string) ([]map[string]types.AttributeValue, error) {
	resp, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("TodoItems"),
		KeyConditionExpression: aws.String("UserID = :userID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userID": &types.AttributeValueMemberS{Value: userID},
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

func UpdateTodoItemStatus(svc *dynamodb.Client, userID, todoID, newStatus string) error {
	_, err := svc.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String("TodoItems"),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: userID},
			"TodoID": &types.AttributeValueMemberS{Value: todoID},
		},
		UpdateExpression: aws.String("SET #status = :newStatus"),
		ExpressionAttributeNames: map[string]string{
			"#status": "Status",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":newStatus": &types.AttributeValueMemberS{Value: newStatus},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	})

	if err != nil {
		return fmt.Errorf("failed to update todo item: %w", err)
	}
	return nil
}

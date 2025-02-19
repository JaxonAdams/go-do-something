package util

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Helper function to extract string attributes safely
func GetStringAttribute(item map[string]types.AttributeValue, key string) string {
	attr, ok := item[key].(*types.AttributeValueMemberS)
	if !ok {
		fmt.Printf("Missing or incorrect type for key: %s\n", key) // Debugging line
		return "N/A"
	}
	return attr.Value
}

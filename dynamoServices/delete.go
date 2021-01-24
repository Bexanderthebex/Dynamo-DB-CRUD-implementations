package dynamoServices

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func DeleteUser(userId int64) *dynamodb.DeleteItemOutput {
	session := NewAWSSession()

	dynamoDBClient := dynamodb.New(session)

	user := GetUser(userId)

	dynamoDBQuery := &dynamodb.DeleteItemInput{
		Key:       user.GetUserKey(),
		TableName: aws.String("user"),
	}

	result, deleteItemError := dynamoDBClient.DeleteItem(dynamoDBQuery)

	if deleteItemError != nil {
		fmt.Println(deleteItemError.Error())
		os.Exit(1)
	}

	fmt.Println(result)

	return result
}

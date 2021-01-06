package dynamoServices

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"sampleDynamo/models"
)

func CreateUser(name string) {
	dynamoUser := models.User{
		Name: name,
	}

	dynamoUser.SetUserId(1)

	session := NewAWSSession()

	dynamoDBClient := dynamodb.New(session)

	dynamoDbUser, mashalMapErr := dynamodbattribute.MarshalMap(dynamoUser)

	if mashalMapErr != nil {
		fmt.Println(mashalMapErr.Error())
	}

	dynamoDBUserRecord := &dynamodb.PutItemInput{
		TableName: aws.String("user"),
		Item:      dynamoDbUser,
	}
	_, err := dynamoDBClient.PutItem(dynamoDBUserRecord)

	if err != nil {
		fmt.Println(err.Error())
	}
}

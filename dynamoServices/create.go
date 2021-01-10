package dynamoServices

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"sampleDynamo/models"
)

func CreateUser(userId int64, name string) {
	dynamoUser := models.User{
		User_id: userId,
		Name:    name,
	}

	session := NewAWSSession()

	dynamoDBClient := dynamodb.New(session)

	dynamoDbUser, mashalMapErr := dynamodbattribute.MarshalMap(dynamoUser)

	if mashalMapErr != nil {
		fmt.Println(mashalMapErr.Error())
		return
	}

	dynamoDBUserRecord := &dynamodb.PutItemInput{
		TableName: aws.String("user"),
		Item:      dynamoDbUser,
	}
	_, err := dynamoDBClient.PutItem(dynamoDBUserRecord)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

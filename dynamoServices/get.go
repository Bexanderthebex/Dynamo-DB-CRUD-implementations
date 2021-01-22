package dynamoServices

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"strconv"

	"sampleDynamo/models"
)

func GetUser(userId int64) models.User {
	session := NewAWSSession()

	dynamoDBClient := dynamodb.New(session)

	dynamoDBQuery := &dynamodb.GetItemInput{
		TableName: aws.String("user"),
		Key: map[string]*dynamodb.AttributeValue{
			"User_id": {
				N: aws.String(strconv.FormatInt(userId, 10)),
			},
		},
	}

	result, error := dynamoDBClient.GetItem(dynamoDBQuery)

	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}

	if result.Item == nil {
		fmt.Printf("Could not find User with Id %d \n", userId)
		os.Exit(1)
	}

	user := models.User{}

	unmarshalUserError := dynamodbattribute.UnmarshalMap(result.Item, &user)

	if unmarshalUserError != nil {
		fmt.Println(unmarshalUserError.Error())
	}

	fmt.Println(user)

	return user
}

package dynamoServices

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
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

func CreateMultipleUsers(users *[]string) {
	usersToCreate := make([]models.User, len(*users))
	for index, userName := range *users {
		userUUID := int64(uuid.New().ID())
		fmt.Println(userUUID, userName)

		usersToCreate[index] = models.User{
			User_id: userUUID,
			Name:    userName,
		}
	}

	session := NewAWSSession()

	dynamoDBClient := dynamodb.New(session)

	createUserInputs := make([]*dynamodb.PutRequest, len(usersToCreate))
	for index, userDocument := range usersToCreate {
		dynamoDbUser, mashalMapErr := dynamodbattribute.MarshalMap(userDocument)

		if mashalMapErr != nil {
			fmt.Println(mashalMapErr.Error())
			return
		}

		createUserInputs[index] = &dynamodb.PutRequest{
			Item: dynamoDbUser,
		}
	}

	userWriteRequests := make([]*dynamodb.WriteRequest, len(createUserInputs))
	for index, userPutRequest := range createUserInputs {
		userWriteRequests[index] = &dynamodb.WriteRequest{
			DeleteRequest: nil,
			PutRequest:    userPutRequest,
		}
	}

	batchWriteItemReq := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			"user": userWriteRequests,
		},
	}

	_, err := dynamoDBClient.BatchWriteItem(batchWriteItemReq)

	if err != nil {
		fmt.Println(err)
		return
	}
}

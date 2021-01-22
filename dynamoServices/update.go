package dynamoServices

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"sampleDynamo/models"
)

func UpdateUser(userId int64, userBalance *models.UserBalance) {
	session := NewAWSSession()

	dynamoDBClient := dynamodb.New(session)

	user := GetUser(userId)

	userBalanceInfo, marshalUserBalanceErr := dynamodbattribute.MarshalMap(userBalance)

	if marshalUserBalanceErr != nil {
		fmt.Println("Got error marshalling the info:")
		fmt.Println(marshalUserBalanceErr.Error())
		os.Exit(1)
	}

	fmt.Println(user.GetUserKey())
	fmt.Println(userBalanceInfo)

	updateUserReq := &dynamodb.UpdateItemInput{
		TableName:                 aws.String("user"),
		ExpressionAttributeValues: userBalanceInfo,
		Key:                       user.GetUserKey(),
		ReturnValues:              aws.String("UPDATED_NEW"),
		UpdateExpression:          aws.String("SET balance = :balance"),
	}

	result, requestErr := dynamoDBClient.UpdateItem(updateUserReq)

	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return
	}

	fmt.Println(result.Attributes)
}

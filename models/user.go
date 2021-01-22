package models

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strconv"
)

type User struct {
	User_id int64
	Name    string
}

type UserBalance struct {
	Balance float64 `json:":balance"`
}

func (u *User) GetUserKey() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"User_id": {
			N: aws.String(strconv.FormatInt(u.User_id, 10)),
		},
	}
}

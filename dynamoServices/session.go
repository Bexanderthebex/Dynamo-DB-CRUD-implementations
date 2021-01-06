package dynamoServices

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSession() *session.Session {
	sess, err := session.NewSessionWithOptions(
		session.Options{
			Profile: "admin-test-user-modal",
			Config:  aws.Config{Region: aws.String("us-west-2")},
		})

	if err != nil {
		fmt.Println(err.Error())
	}

	return sess
}

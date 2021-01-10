package main

import "sampleDynamo/dynamoServices"

func main() {
	dynamoServices.CreateUser(1, "best boi")

	dynamoServices.GetUser(1)
}

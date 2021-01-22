package main

import "sampleDynamo/dynamoServices"
import "sampleDynamo/models"

func main() {
	dynamoServices.CreateUser(1, "best boi")

	dynamoServices.GetUser(1)

	batchOfUsersToCreate := []string{"etoro", "robinhood", "TD ameritrade", "JPM", "Modal"}

	dynamoServices.CreateMultipleUsers(&batchOfUsersToCreate)

	dynamoServices.UpdateUser(1, &models.UserBalance{Balance: 11})
}

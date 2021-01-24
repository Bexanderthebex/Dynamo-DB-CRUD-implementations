package main

import "sampleDynamo/dynamoServices"

func main() {
	//dynamoServices.CreateUser(1, "best boi")
	//
	//dynamoServices.GetUser(1)
	//
	//batchOfUsersToCreate := []string{"etoro", "robinhood", "TD ameritrade", "JPM", "Modal"}
	//
	//dynamoServices.CreateMultipleUsers(&batchOfUsersToCreate)
	//
	//dynamoServices.UpdateUser(1, &models.UserBalance{Balance: 11})

	dynamoServices.DeleteUser(1)
}

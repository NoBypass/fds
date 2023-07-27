package main

import (
	"fmt"
	"server/db/models"
	"server/db/repository"
	dbutils "server/db/utils"
)

func main() {
	fmt.Println("Starting server...")

	driver, ctx, _ := dbutils.ConnectDB()

	fmt.Println("Connected to database")

	fmt.Println("Server is running on port 8080")

	account, err := repository.CreateAccount(ctx, driver, &models.Account{
		ID:       "1",
		Username: "test",
		Password: "test",
		JoinedAt: 0,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(account)
	
	dbutils.CloseDB(driver, ctx)
}

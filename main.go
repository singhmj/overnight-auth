package main

import (
	"auth/api"
	"auth/db"
	"auth/utils"
	"fmt"
)

// TODO: convert this to a struct
var (
	DBHost = "localhost"
	DBPort = 5432
	DBUser = "postgres"
	DBPass = "postgres"
	DBName = "auth"
)

func main() {
	// TODO: graceful shutdown
	// TODO: read from env variables or config file

	dbConn, err := utils.NewDB("localhost", DBPort, DBUser, DBPass, DBName)
	if err != nil {
		panic(fmt.Errorf("an error occured while connecting with database, more: %v", err))
	}

	store := db.NewStore(dbConn)
	server := api.NewServer(&store)

	fmt.Println(server.Start("localhost:9090"))
}

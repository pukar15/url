package main

import (
	"github.com/pukarlamichhane/url-shorter.git/database"
	"github.com/pukarlamichhane/url-shorter.git/router"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database.ConnectDb()
}

func main() {
	router.ClientRoutes()
}

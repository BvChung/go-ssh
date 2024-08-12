package main

import (
	"log"

	"github.com/BvChung/go-ssh/cmd/ssh/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}

	// server, err := models.CreateServer()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db.GetDatabaseCredentials("new-db")
	db.CreateDatabaseToken("new-db")
	// server.Start()
	// db.CreateDB()
	// db.ListDB()
	// db.ValidateToken()
}

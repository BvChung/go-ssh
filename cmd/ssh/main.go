package main

import (
	"log"

	"github.com/BvChung/go-ssh/cmd/ssh/serve"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}

	server, err := serve.CreateServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Start()

	// db.GetDatabaseCredentials("new-db")
	// db.CreateDatabaseToken("new-db")
	// db.CreateDB()
	// db.ListDB()
	// db.ValidateToken()
}

package main

import "github.com/BvChung/go-ssh/cmd/ssh/db"

func main() {
	// server, err := models.CreateApp()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db.GetDatabaseCredentials("new-db")
	// server.Start()
	// db.CreateDB()
	// db.ListDB()
	// db.ValidateToken()
}

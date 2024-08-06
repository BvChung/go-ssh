package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseCreationArgs struct {
	Name  string `json:"name"`
	Group string `json:"group"`
}

func NewDatabaseCreationArgs(name, group string) DatabaseCreationArgs {
	return DatabaseCreationArgs{Name: name, Group: group}
}

func CreateDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}

	token := os.Getenv("API_TOKEN")
	org := os.Getenv("ORG_NAME")

	args, err := json.Marshal(NewDatabaseCreationArgs("new-db", "default"))

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", org), bytes.NewBuffer(args))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(res.Status)
	fmt.Println(string(data))
}

type FindDatabaseCredentials struct {
	token string

}

func GetDatabaseCredentials(dbName string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}

	token := os.Getenv("API_TOKEN")
	org := os.Getenv("ORG_NAME")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s", org, dbName), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	fmt.Println(string(data))
}

func ListDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}

	token := os.Getenv("API_TOKEN")
	org := os.Getenv("ORG_NAME")

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", org), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	data, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

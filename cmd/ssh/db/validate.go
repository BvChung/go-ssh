package db

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ListDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}

	token := os.Getenv("API_TOKEN")
	org := os.Getenv("ORG_NAME")

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases", org), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Status)
	data, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

}

func ValidateToken() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}

	token := os.Getenv("API_TOKEN")

	req, err := http.NewRequest("GET", "https://api.turso.tech/v1/auth/validate", nil)

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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Unable to read res data")

	}

	fmt.Println(res.Status)
	fmt.Println(string(data))
}

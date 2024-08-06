package db

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ValidateToken() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}

	token := os.Getenv("API_TOKEN")

	req, err := http.NewRequest(http.MethodGet, "https://api.turso.tech/v1/auth/validate", nil)

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

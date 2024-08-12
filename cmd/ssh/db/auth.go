package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	tealog "github.com/charmbracelet/log"
)

type DatabaseToken struct {
	Jwt string `json:"jwt"`
}

func CreateDatabaseToken(databaseName string) (string, error) {
	/*
				curl -L -X POST 'https://api.turso.tech/v1/organizations/{organizationName}/databases/{databaseName}/auth/tokens?expiration=2w&authorization=full-access' \
		  		-H 'Authorization: Bearer TOKEN'
	*/
	token := os.Getenv("API_TOKEN")
	org := os.Getenv("ORG_NAME")

	endpoint := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/databases/%s/auth/tokens?expiration=2w&authorization=full-access", org, databaseName)
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var dbToken DatabaseToken

	if err = json.NewDecoder(res.Body).Decode(&dbToken); err != nil {
		return "", err
	}

	tealog.Info(dbToken)

	return "", nil
}

func ValidateToken() {
	token := os.Getenv("API_TOKEN")

	req, err := http.NewRequest(http.MethodGet, "https://api.turso.tech/v1/auth/validate", nil)

	if err != nil {
		tealog.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		tealog.Fatal(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		tealog.Fatal("Unable to read res data")

	}

	fmt.Println(res.Status)
	fmt.Println(string(data))
}

package helperFunc

import (
	"os"

	"github.com/joho/godotenv"
)

type ApiCredentials struct {
	ID       string
	Secret   string
	Username string
	Password string
}

func GetApiCredentials() (apiCredentials ApiCredentials, err error) {
	if err := godotenv.Load("apiCredentials.env"); err != nil {
		return ApiCredentials{}, err
	}

	apiCredentials = ApiCredentials{}

	apiCredentials.ID = os.Getenv("ID")
	apiCredentials.Secret = os.Getenv("SECRET")
	apiCredentials.Password = os.Getenv("PASSWORD")

	return apiCredentials, nil

}

package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	TWILIO_ACCOUNT_SID_KEY = "TWILIO_ACCOUNT_SID"
	TWILIO_AUTHTOKEN_KEY   = "TWILIO_AUTHTOKEN"
	TWILIO_SERVICES_ID_KEY = "TWILIO_SERVICES_ID"
	VONAGE_API_KEY         = "VONAGE_API_KEY"
	VONAGE_API_SECRET      = "VONAGE_API_SECRET"
)

type TwilioConfig struct {
	TWILIO_ACCOUNT_SID string
	TWILIO_AUTHTOKEN   string
	TWILIO_SERVICES_ID string
}

type VonageConfig struct {
	ApiKey    string
	ApiSecret string
}

func LoadConfig() (vonageConfig *VonageConfig, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Load env file fail with %v", err)
		return nil, err
	}
	ApiKey := os.Getenv(VONAGE_API_KEY)
	ApiSecret := os.Getenv(VONAGE_API_SECRET)
	vonageConfig = &VonageConfig{
		ApiKey:    ApiKey,
		ApiSecret: ApiSecret,
	}
	return vonageConfig, nil
}

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
)

type TwilioConfig struct {
	TWILIO_ACCOUNT_SID string
	TWILIO_AUTHTOKEN   string
	TWILIO_SERVICES_ID string
}

func loadENV() (twilioConfig *TwilioConfig, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Load env file fail with %v", err)
		return nil, err
	}
	TWILIO_ACCOUNT_SID := os.Getenv(TWILIO_ACCOUNT_SID_KEY)
	TWILIO_AUTHTOKEN := os.Getenv(TWILIO_AUTHTOKEN_KEY)
	TWILIO_SERVICES_ID := os.Getenv(TWILIO_SERVICES_ID_KEY)
	twilioConfig = &TwilioConfig{
		TWILIO_ACCOUNT_SID: TWILIO_ACCOUNT_SID,
		TWILIO_AUTHTOKEN:   TWILIO_AUTHTOKEN,
		TWILIO_SERVICES_ID: TWILIO_SERVICES_ID,
	}
	return twilioConfig, nil
}

package api

import (
	"fmt"

	"github.com/dchest/uniuri"
	"github.com/vonage/vonage-go-sdk"
)

type SMSService interface {
	SendSMS(phoneNumber string) error
	VerifySMS(phoneNumber string, code string) error
}

type VonageSMSService struct {
	Config *VonageConfig
}

func NewVonageSMSService(config *VonageConfig) *VonageSMSService {
	return &VonageSMSService{
		Config: config,
	}
}

func (service *VonageSMSService) SendSMS(phoneNumber string) error {
	auth := vonage.CreateAuthFromKeySecret(service.Config.ApiKey, service.Config.ApiSecret)
	smsClient := vonage.NewSMSClient(auth)
	otp := uniuri.NewLen(6)
	response, errResp, err := smsClient.Send("Vonage APIs", phoneNumber, fmt.Sprintf("OTP: %s", otp), vonage.SMSOpts{})
	if err != nil {
		return err
	}
	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	} else {
		fmt.Println("Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
	}
	return nil
}

func (service *VonageSMSService) VerifySMS(phoneNumber string, code string) error {
	return nil
}

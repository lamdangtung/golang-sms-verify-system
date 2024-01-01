package main

import (
	"log"

	"com.lamdt/sms-verify-system/api"
	"github.com/gin-gonic/gin"
)

func main() {

	config, err := api.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	smsService := api.NewVonageSMSService(config)

	router := gin.Default()
	server := api.NewServer(router, smsService)

	server.Routes()

	router.Run(":9999")
}

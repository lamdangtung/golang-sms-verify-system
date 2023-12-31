package main

import (
	"com.lamdt/sms-verify-system/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server := api.NewServer(router)

	server.Routes()

	router.Run(":9999")
}

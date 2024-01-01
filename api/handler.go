package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendSMSRequest struct {
	PhoneNumber string `json:"phone_number", binding:"required,e164"`
}

func (server *Server) sendSMS(context *gin.Context) {
	var req SendSMSRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	phoneNumber := req.PhoneNumber
	server.SmsService.SendSMS(phoneNumber)
	context.JSON(http.StatusAccepted, gin.H{"status": "success", "message": "OTP sent successfully"})
}

func (server *Server) verifySMS(context *gin.Context) {

}

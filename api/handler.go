package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendSMSRequest struct {
	PhoneNumber string `json: "phone_number", binding:"required"`
}

func (server *Server) sendSMS(context *gin.Context) {
	var req SendSMSRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"phone_number": req.PhoneNumber})
}

func (server *Server) verifySMS(context *gin.Context) {

}

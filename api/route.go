package api

import "github.com/gin-gonic/gin"

type Server struct {
	Router     *gin.Engine
	SmsService SMSService
}

func NewServer(router *gin.Engine, smsService SMSService) *Server {
	return &Server{
		Router:     router,
		SmsService: smsService,
	}
}

func (server *Server) Routes() {
	server.Router.POST("/otp", server.sendSMS)
	server.Router.POST("/otp/verify", server.verifySMS)
}

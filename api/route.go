package api

import "github.com/gin-gonic/gin"

type Server struct {
	Router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{
		Router: router,
	}
}

func (server *Server) Routes() {
	server.Router.POST("/otp", server.sendSMS)
	server.Router.POST("/otp/verify", server.verifySMS)
}

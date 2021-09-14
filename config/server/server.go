package server

import (
	"log"

	"estudar.com.br/config/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "estudar.com.br/docs"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.New(),
	}
}

func (ser *Server) Run() {
	router := routes.ConfigRoutes(ser.server)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Print("Server rodando porta" + ser.port)
	log.Fatal(router.Run(":" + ser.port))
}

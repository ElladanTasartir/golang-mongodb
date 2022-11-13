package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *gin.Engine
}

func Run(port int) *Server {
	server := Server{
		httpServer: gin.Default(),
	}

	fmt.Println(fmt.Sprintf("Successfully started application on port %d", port))

	return &server
}

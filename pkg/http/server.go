package http

import (
	"fmt"
	"net/http"

	"github.com/ElladanTasartir/golang-mongodb/pkg/storage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *gin.Engine
	client     *storage.DbClient
}

func Run(port int, client *storage.DbClient) *Server {
	server := Server{
		httpServer: gin.New(),
		client:     client,
	}

	server.httpServer.Use(gin.Logger())

	server.AddDecksEndpoints()
	server.httpServer.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "All done",
		})
	})

	server.httpServer.Run(fmt.Sprintf(":%d", port))

	fmt.Println(fmt.Sprintf("Successfully started application on port %d", port))

	return &server
}

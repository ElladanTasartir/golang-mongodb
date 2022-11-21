package http

import "github.com/gin-gonic/gin"

type Response struct {
	Code int
	Body interface{}
}

func SendResponse(ctx *gin.Context, response *Response) {
	ctx.JSON(response.Code, response.Body)
}

package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器出错",
	})
}

func RespErrorWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "false",
		"data":   data,
	})
}

func RespSuccessfulWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "true",
		"data":   data,
	})
}

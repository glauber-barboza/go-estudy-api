package httputil

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func ExpectedExceptionErr(ctx *gin.Context, err error) {
	log.Println(err.Error())
	er := HTTPError{
		Code:    http.StatusExpectationFailed,
		Message: err.Error(),
	}
	ctx.JSON(er.Code, er)
}

func ExpectedException(ctx *gin.Context, message string, err error) {
	log.Println(message, err.Error())
	er := HTTPError{
		Code:    http.StatusExpectationFailed,
		Message: message,
	}
	ctx.JSON(er.Code, er)
}

func HttpException(ctx *gin.Context, status int, message string, err error) {
	log.Println(message, err.Error())
	er := HTTPError{
		Code:    status,
		Message: message,
	}
	ctx.JSON(er.Code, er)
}

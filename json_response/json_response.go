package json_response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, APIResponse{
		Code:    "00",
		Message: "success",
		Data:    data,
	})
	panic(nil)
}

func Error(ctx *gin.Context, httpCode int, code string, message string, err interface{}) {
	ctx.JSON(httpCode, APIResponse{
		Code:    code,
		Message: message,
		Data:    "",
		Error:   err,
	})
	panic(nil)
}

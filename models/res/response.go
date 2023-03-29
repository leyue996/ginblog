package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success = 0
	Error   = 7
)

type Response struct {
	code int
	data any
	msg  string
}

func Result(code int, data any, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		code: code,
		data: data,
		msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "成功", c)
}
func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(int(code), map[string]any{}, "未知错误", c)
}

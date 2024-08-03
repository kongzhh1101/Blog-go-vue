package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Error   = 7
	Success = 0
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ResponseList[T any] struct {
	Count any `json:"count"`
	List  []T `json:"list"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(Success, nil, msg, c)
}

func OKWithList[T any](count any, list []T, c *gin.Context) {
	data := ResponseList[T]{
		Count: count,
		List:  list,
	}
	Result(Success, data, "操作成功", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, nil, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), nil, msg, c)
		return
	}
	Result(Error, nil, "未知错误", c)
}

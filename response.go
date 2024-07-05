package x14nfile

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess      = 0
	CodeFail         = 1
	CodeUnknownError = -1
	CodeSessionError = 40000
)

const (
	YYYYMMDD       = "2006-01-02"
	YYYYMMDDHHIISS = "2006-01-02 15:04:05"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
	Time string `json:"time"`
}

func ResultJson(ctx *gin.Context, code int, msg string, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
		Time: time.Now().Format(YYYYMMDDHHIISS),
	})
}

func Success(ctx *gin.Context) {
	ResultJson(ctx, CodeSuccess, "success", nil)
}

func SuccessWithData(ctx *gin.Context, data any) {
	ResultJson(ctx, CodeSuccess, "success", data)
}

func Fail(ctx *gin.Context) {
	ResultJson(ctx, CodeFail, "fail", nil)
}

func FailWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, CodeFail, msg, nil)
}

package apiwrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Code int    `json:"code,omitempty"`
	Data T      `json:"data,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func respond[T any](code int, data T, msg string) *Response[T] {
	return &Response[T]{Code: code, Data: data, Msg: msg}
}

func Success[T any](data T) *Response[T] {
	return respond(http.StatusOK, data, "操作成功")
}

func SuccessWithMsg[T any](data T, msg string) *Response[T] {
	return respond(http.StatusOK, data, msg)
}

func Fail(code int) *Response[any] {
	return respond[any](code, nil, "操作失败")
}

func FailWithMsg(code int, msg string) *Response[any] {
	return respond[any](code, nil, msg)
}

func Wrap[T any](fn func(ctx *gin.Context) (T, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := fn(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, data)
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}

func WrapWithBody[R any, T any](fn func(ctx *gin.Context, req R) (T, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req R
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, FailWithMsg(http.StatusBadRequest, err.Error()))
			return
		}
		data, err := fn(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusOK, data)
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}

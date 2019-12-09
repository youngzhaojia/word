package app

import (
	"github.com/gin-gonic/gin"
	"word/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(code int, ret int, msg string, data interface{}) {
	g.C.JSON(code, Response{
		Ret:  ret,
		Msg:  msg,
		Data: data,
	})
	return
}

func (g *Gin) ResponseErr(ret int, msg string) {
	g.Response(200, ret, msg, nil)
	return
}

func (g *Gin) ResponseErrMsg(msg string) {
	g.ResponseErr(e.ERROR, msg)
	return
}

func (g *Gin) ResponseSuccess(msg string, data interface{}) {
	g.Response(200, e.SUCCESS, msg, data)
	return
}

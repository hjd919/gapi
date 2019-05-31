package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hjd919/gapi/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) Res(data interface{}, code int) {
	msg := ""
	if code != 200 {
		msg = data.(string)
	}

	g.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}

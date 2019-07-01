package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ResponseData struct {
	Err  int 			`json:"err"`
	Msg  string 		`json:"msg"`
	Data interface{} 	`json:"data,omitempty"`
}

func RespondToClient(c echo.Context, errCode int, msg string, data interface{}) error {

	var response ResponseData

	response.Err = errCode
	response.Msg = msg
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

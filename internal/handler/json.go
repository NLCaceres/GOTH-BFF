package handler

import (
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/labstack/echo/v4"
	"net/http"
)

func QueryJSON(c echo.Context) error {
	res, err := queryapi.Call(c.Path()[1:])
	if err != nil {
		return c.NoContent(queryapi.ErrCode(err))
	}
	return c.JSON(http.StatusOK, queryResponse{res.Documents()})
}

type queryResponse struct {
	Results []queryapi.Document
}

package handler

import (
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/labstack/echo/v4"
	"net/http"
)

func QueryJSON(c echo.Context) error {
	res, err := queryapi.Call(c.Path()[1:])
	if e, ok := err.(queryapi.Error); ok {
		return c.NoContent(e.Code)
	}
	return c.JSON(http.StatusOK, queryResponse{res.Documents()})
}

type queryResponse struct {
	Results []queryapi.Document
}

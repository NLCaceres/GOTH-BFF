package handler

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/util/url"
	"github.com/labstack/echo/v5"
	"net/http"
)

func QueryJSON(c *echo.Context) error {
	res, err := queryapi.Call(url.New(*c.Request().URL))
	var e queryapi.Error
	if errors.As(err, &e) {
		return c.NoContent(e.Code)
	}
	return c.JSON(http.StatusOK, queryResponse{res.Documents()})
}

type queryResponse struct {
	Results []queryapi.Document
}

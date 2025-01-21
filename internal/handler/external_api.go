package handler

import (
	"bytes"
	"encoding/json"
	"github.com/NLCaceres/goth-example/internal/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
)

// POSTs pre-formatted JSON to an API after dynamically updating the JSON string's
// key-value pair corresponding to the search value
func ApiPostRequest(c echo.Context) error {
	queryMap, err := util.ReadJSON[map[string][]map[string]any]("internal/query.json")
	if err != nil {
		log.Errorf("Issue getting formatted JSON query map due to: %s\n", err)
	}

	queries := queryMap["searches"]
	queries[len(queries)-1]["q"] = c.Path()[1:] // Drop 1st "/". No Unicode in URLs so OK

	jsonBytes, err := json.MarshalIndent(queryMap, "", "  ") // Last param sets spacing
	if err != nil {
		log.Errorf("Issue parsing JSON map into a []byte due to: %s\n", err)
	}

	response, err := util.PostJSON(os.Getenv("EXTERNAL_API_URL"), bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Errorf("Issue making POST Request due to: %s\n", err)
	}

	return c.JSON(http.StatusOK, response)
}

package handler

import (
	"bytes"
	"encoding/json"
	"github.com/NLCaceres/goth-example/internal/util"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

// POSTs pre-formatted JSON to an API after dynamically updating the JSON string's
// key-value pair corresponding to the search value
func ApiPostRequest(c echo.Context) error {
	queryMap, err := util.ReadJSON[map[string][]map[string]any](os.Getenv("QUERY_FILE"))
	if err != nil {
		log.Printf("Issue getting formatted JSON query map due to: %s\n", err)
		return c.NoContent(500)
	}

	queries := queryMap["searches"]
	queries[len(queries)-1]["q"] = c.Path()[1:] // Drop 1st "/". No Unicode in URLs so OK

	jsonBytes, err := json.MarshalIndent(queryMap, "", "  ")
	if err != nil { // Unclear if Marshal can even fail since it parses already parsed JSON
		log.Printf("Issue parsing JSON map into a []byte due to: %s\n", err)
		return c.NoContent(400)
	}

	response, err := util.PostJSON(os.Getenv("EXTERNAL_API_URL"), bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Printf("Issue making POST Request due to: %s\n", err)
		return c.NoContent(502)
	}

	return c.JSON(http.StatusOK, response)
}

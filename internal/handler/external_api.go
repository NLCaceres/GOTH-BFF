package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"strings"
)

// POSTs pre-formatted JSON to an API after dynamically updating the JSON string's
// key-value pair corresponding to the search value
func ApiPostRequest(c echo.Context) error {
	queryMap, err := util.ReadJSON[map[string][]map[string]any](os.Getenv("QUERY_FILE"))
	if err != nil {
		log.Printf("Issue getting formatted JSON query map due to: %s\n", err)
		return c.NoContent(500) // Internal issue
	}

	queries := queryMap["searches"]
	queries[len(queries)-1]["q"] = c.Path()[1:] // Drop 1st "/". No Unicode in URLs so OK
	if err := setFilters(queries[len(queries)-1]); err != nil {
		log.Print("Issue setting filters due to:", err)
		return c.NoContent(501) // Implementation issue
	}

	jsonBytes, err := json.MarshalIndent(queryMap, "", "  ")
	if err != nil { // Unclear if Marshal can even fail since it parses already parsed JSON
		log.Printf("Issue parsing JSON map into a []byte due to: %s\n", err)
		return c.NoContent(400) // Bad request probably due to changes in JSON map
	}

	response, err := util.PostJSON(os.Getenv("EXTERNAL_API_URL"), bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Printf("Issue making POST Request due to: %s\n", err)
		return c.NoContent(502) // Gateway error due to upstream server issue
	}

	return c.JSON(http.StatusOK, response)
}

func setFilters(jsonObj map[string]any) error {
	filter, ok := jsonObj["filter_by"].(string) // Type coercion
	if !ok {
		errMsg := fmt.Sprintf("Issue coercing JSON filter value %v to string", jsonObj["filter_by"])
		return errors.New(errMsg)
	}
	matches, err := util.FindDunderVars(filter)
	if err != nil {
		return err
	}

	replacements := strings.Split(os.Getenv("FILTER_REPLACEMENTS"), "|")
	length := len(replacements)
	if len(matches) < length { //NOTE: These few lines are basically how Python's Zip works
		length = len(matches)
	}
	for i := 0; i < length; i++ { // Matching all values together until one array runs out
		filter = strings.Replace(filter, matches[i], replacements[i], 1)
	}
	jsonObj["filter_by"] = filter
	return nil
}

package queryapi

import (
	"bytes"
	"encoding/json"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/proxy"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

// POSTs pre-formatted JSON to an API after dynamically updating the JSON string's
// key-value pair corresponding to the search value
func NewQuery(c echo.Context) error {
	queryReq, err := fileread.JSON[Request](os.Getenv("QUERY_FILE"))
	if err != nil {
		log.Printf("Issue getting formatted JSON query map due to: %s\n", err)
		return c.NoContent(500) // Internal issue
	}

	queryReq.Terms[len(queryReq.Terms)-1].Q = c.Path()[1:] // Drop 1st "/"
	if err := queryReq.Terms[len(queryReq.Terms)-1].setFilters(os.Getenv("FILTER_REPLACEMENTS")); err != nil {
		log.Print("Issue setting filters due to:", err)
		return c.NoContent(501) // Implementation issue
	}

	jsonBytes, err := json.MarshalIndent(queryReq, "", "  ")
	if err != nil { // Unclear if Marshal can even fail since it parses already parsed JSON
		log.Printf("Issue parsing JSON map into a []byte due to: %s\n", err)
		return c.NoContent(400) // Bad request probably due to changes in JSON map
	}

	res, err := proxy.PostJSON(os.Getenv("EXTERNAL_API_URL"), bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Printf("Issue making POST Request due to: %s\n", err)
		return c.NoContent(502) // Gateway error due to upstream server issue
	}

	return c.JSON(http.StatusOK, res)
}

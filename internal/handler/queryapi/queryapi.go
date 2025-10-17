package queryapi

import (
	"bytes"
	"encoding/json"
	"errors"
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
	query, err := newQuery(os.Getenv("QUERY_FILE"))
	if err != nil {
		return c.NoContent(queryErrCode(err))
	}
	res, err := proxy.PostJSON(os.Getenv("EXTERNAL_API_URL"), query)
	if err != nil {
		log.Printf("Issue making POST Request due to: %s\n", err)
		return c.NoContent(502) // Gateway error due to upstream server issue
	}
	return c.JSON(http.StatusOK, res)
}

func newQuery(path string) (*bytes.Buffer, error) {
	queryReq, err := fileread.JSON[Request](os.Getenv("QUERY_FILE"))
	if err != nil {
		log.Printf("Issue getting formatted JSON query map due to: %s\n", err)
		return nil, err
	}

	queryReq.last().Q = path
	if err := queryReq.last().setFilters(os.Getenv("FILTER_REPLACEMENTS")); err != nil {
		log.Print("Issue setting filters due to:", err)
		return nil, SearchSetterError
	}

	jsonBytes, err := json.MarshalIndent(queryReq, "", "  ")
	if err != nil { // CAN get a `jsonErr` or `scanErr` if `Q` is poisoned BUT unclear how
		log.Printf("Issue parsing JSON map into a []byte due to: %s\n", err)
		return nil, err
	}
	return bytes.NewBuffer(jsonBytes), nil
}

func queryErrCode(e error) int {
	if errors.As(e, new(fileread.FileReadError)) {
		return 500 // Internal issue
	} else if errors.Is(e, SearchSetterError) {
		return 501 // Implementation issue
	} else {
		return 400 // Bad request
	}
}

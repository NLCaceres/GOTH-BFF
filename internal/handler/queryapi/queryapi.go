package queryapi

import (
	"bytes"
	"encoding/json"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/proxy"
	appUrl "github.com/NLCaceres/goth-example/internal/util/url"
	"log"
	"net/url"
	"os"
	"strconv"
)

// POSTs pre-formatted JSON to an API after dynamically updating the JSON string's
// key-value pair corresponding to the search value
func Call(URL url.URL) (*Response, error) {
	pageQuery := URL.Query().Get("page")
	if pageQuery == "" {
		pageQuery = "1"
	}
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		log.Printf("Page %v converted to int %d failed: %v", pageQuery, page, err)
		return nil, NewError(appUrl.ParamError{Key: "page", Value: pageQuery})
	}
	query, err := newQuery(URL.Path[1:], page)
	if err != nil {
		return nil, NewError(err)
	}
	res, err := proxy.PostJSON[*Response](os.Getenv("EXTERNAL_API_URL"), query)
	if err != nil {
		log.Printf("Issue making POST Request due to: %s\n", err)
		return nil, NewError(err)
	}
	return res, nil
}

func newQuery(path string, page int) (*bytes.Buffer, error) {
	queryReq, err := fileread.JSON[Request](os.Getenv("QUERY_FILE"))
	if err != nil {
		log.Printf("Issue getting formatted JSON query map due to: %s\n", err)
		return nil, err
	}

	queryReq.last().Page = page
	if page < 1 {
		log.Printf("Got page < 1 equal to %d", page)
		return nil, SearchPageError
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

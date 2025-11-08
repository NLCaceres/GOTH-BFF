package queryapi

import (
	"bytes"
	"encoding/json"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/proxy"
	"log"
	"os"
)

// POSTs pre-formatted JSON to an API after dynamically updating the JSON string's
// key-value pair corresponding to the search value
func Call(path string) (*Response, error) {
	query, err := newQuery(path)
	if err != nil {
		return nil, err
	}
	res, err := proxy.PostJSON[*Response](os.Getenv("EXTERNAL_API_URL"), query)
	if err != nil {
		log.Printf("Issue making POST Request due to: %s\n", err)
		return nil, err
	}
	return res, nil
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

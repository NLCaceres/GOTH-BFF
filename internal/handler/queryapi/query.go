package queryapi

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/util/datetime"
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"strings"
)

type Response struct {
	Results []struct {
		Hits []struct {
			Document Document `json:"document"`
		} `json:"hits"`
	} `json:"results"`
}
type Document struct {
	PostTime datetime.Local `json:"posted"` // Tends to get microsecond precision
	Title    string         `json:"title"`
	URL      string         `json:"url"`
}

type JSONValueString string

func (n *JSONValueString) UnmarshalJSON(b []byte) error {
	*n = JSONValueString(string(b))
	return nil
}

type Search struct {
	FilterBy string `json:"filter_by"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	Q        string `json:"q"`
	SelectBy string `json:"query_by"`
	SortBy   string `json:"sort_by"`
}

var SearchSetterError = errors.New("Unable to set Search Filter Field")

func (s *Search) setFilters(filters string) error {
	matches, err := stringy.FindDunderVars(s.FilterBy)
	if err != nil {
		return SearchSetterError
	}
	replacements := strings.Split(filters, "|")
	for i := range min(len(replacements), len(matches)) {
		s.FilterBy = strings.Replace(s.FilterBy, matches[i], replacements[i], 1)
	}
	return nil
}

type Request struct {
	Terms []Search `json:"searches"`
}

func (r Request) last() *Search {
	if len(r.Terms) == 0 {
		return nil
	}
	return &r.Terms[len(r.Terms)-1]
}

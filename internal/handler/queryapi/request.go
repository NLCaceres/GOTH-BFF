package queryapi

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"strings"
)

type Request struct {
	Terms []Search `json:"searches"`
}

func (r Request) last() *Search {
	if len(r.Terms) == 0 {
		return nil
	}
	return &r.Terms[len(r.Terms)-1]
}

type Search struct {
	Collection string `json:"collection"`
	FilterBy   string `json:"filter_by"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Q          string `json:"q"`
	SelectBy   string `json:"query_by"`
	SortBy     string `json:"sort_by"`
}

var SearchPageError = errors.New("Search Page less than 1")
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

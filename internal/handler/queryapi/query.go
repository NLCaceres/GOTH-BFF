package queryapi

import (
	"github.com/NLCaceres/goth-example/internal/util/datetime"
)

type Response struct {
	Results []struct {
		Hits []struct {
			Document Document `json:"document"`
		} `json:"hits"`
	} `json:"results"`
	documents []Document
}

func NewResponse(docs []Document) *Response {
	res := Response{
		[]struct {
			Hits []struct {
				Document Document `json:"document"`
			} `json:"hits"`
		}{{[]struct {
			Document Document `json:"document"`
		}{},
		}}, nil,
	}
	res.documents = docs
	for _, doc := range docs {
		res.Results[0].Hits = append(res.Results[0].Hits, struct {
			Document Document `json:"document"`
		}{doc})
	}
	return &res
}

func (r *Response) Documents() []Document {
	if len(r.documents) != 0 { // Handles nil since nil slices are length 0
		return r.documents
	}
	for _, result := range r.Results {
		for _, hit := range result.Hits {
			r.documents = append(r.documents, hit.Document)
		}
	}
	return r.documents
}
func (r *Response) SetDocuments(docs []Document) {
	if len(r.Results[0].Hits) > 0 || len(r.documents) > 0 {
		r.Results[0].Hits = r.Results[0].Hits[:0]
		r.documents = r.documents[:0]
	}
	for _, doc := range docs {
		r.Results[0].Hits = append(r.Results[0].Hits, struct {
			Document Document `json:"document"`
		}{doc})
	}
	r.documents = append(r.documents, docs...)
}

type Document struct {
	PostTime datetime.Local `json:"posted"` // Tends to get microsecond precision
	Title    string         `json:"title"`
	URL      string         `json:"url"`
}

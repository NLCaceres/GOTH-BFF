package queryapi

type Search struct {
	FilterBy string `json:"filter_by"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	Q        string `json:"q"`
	SelectBy string `json:"query_by"`
	SortBy   string `json:"sort_by"`
}

type Request struct {
	Terms []Search `json:"searches"`
}

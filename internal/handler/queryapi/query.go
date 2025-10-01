package queryapi

type Search struct {
	FilterBy string
	Page     int
	PerPage  int
	Q        string
	SelectBy string
	SortBy   string
}

type Request struct {
	Terms []Search
}

package model

type Item struct {
	Name        string
	URL         string
	Description string
}

func MockItems() []Item {
	return []Item{
		{Name: "Fo", URL: "Ba", Description: "Lorem ipsum dolor sit amet"},
		{Name: "Fi", URL: "Bu", Description: ""},
		{
			Name: "Za", URL: "Ze",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusm",
		},
	}
}

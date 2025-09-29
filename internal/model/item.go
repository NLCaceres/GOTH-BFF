package model

import "github.com/NLCaceres/goth-example/internal/util/slice"

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
		{Name: "Do", URL: "Re", Description: "Lorem ipsum dolor sit amet, consectetur elit a"},
		{Name: "Mi", URL: "Fa", Description: "Lorem ipsum dolor sit amet, sed do eiusmod"},
		{
			Name: "So", URL: "La",
			Description: "Lorem ipsum dolor sit amet, sed do eiusmod tempor incididunt",
		},
	}
}

func ManyMockItems() []Item {
	return slice.Duplicate(MockItems(), 6)
}

package htmx

import (
	"context"
	"github.com/NLCaceres/goth-example/internal/view/reusable/htmx"
	"github.com/a-h/templ"
)

// Wrapper for templ.Component to add helpful methods to more easily build HTMX component responses
type PageComponent struct {
	templ.Component
}

// Initializer for PageComponent to wrap a templ.Component and deliver as HTMX payload
func Data(c templ.Component) PageComponent {
	return PageComponent{c}
}

// Fluent API to templ.Join the main wrapped component with a Title element so HTMX
// can update the HTML <title> via auto-magic out-of-band swap
func (c PageComponent) AddTitle(title string) PageComponent {
	return PageComponent{templ.Join(c, htmx.Title(title))}
}

// Fluent API to templ.Join the main wrapped component with a Style element so HTMX
// can update via out-of-band swap any matching target, ideally in the <head>
func (c PageComponent) AddStyle(path string) PageComponent {
	return PageComponent{templ.Join(c, htmx.StyleLink(path))}
}

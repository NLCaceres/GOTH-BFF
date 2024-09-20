# GOTH Example - Go + Templ + HTMX
An all around learning experience where I intend to explore Go in greater depths
as well as evaluate the many tools available in the Go ecosystem, all while I get
a better grasp on NeoVim as an IDE, personalizing it to my needs.

### Future Plans
- Make each branch the next major step in development
1. [Echo framework](https://github.com/labstack/echo) to set up the page routing
2. [Templ layout templates](https://github.com/a-h/templ) to set up basic HTML structure of each page
3. Add in [HTMX](https://htmx.org/docs/#introduction) to allow partial page updates in relation to Nav bar or tab bar
4. Setup Search bar that makes POST requests to a GraphQL API of your choice and lists results
5. Swap out Echo for a simpler [Chi implementation](https://github.com/go-chi/chi) of the page routing
6. Swap out Chi for an even simpler [standard library implementation](https://pkg.go.dev/net/http#ServeMux)
   - [Recent improvements to Go's router](https://go.dev/blog/routing-enhancements)
   - [Advanced patterns in Go standard library routing video](https://www.youtube.com/watch?v=H7tbjKFSg58)
7. Dockerize for easy building, sharing and running

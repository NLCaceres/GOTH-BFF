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

### Testing
Golang testing is pretty awesome! It is particularly awesome, because everything you need is baked into Go's
standard library! Testing files end with `_test.go`, and all test functions start with `Test` and include a
parameter `t` of type `*testing.T` (or sometimes of type `*testing.B` or `*testing.F` for Benchmarks or Fuzzing respectively).

For funcs with lots of cases to test, Golang uses table testing as a means of covering all of these test
cases from a single parent test func, structuring the test to be very simple, readable and quick to write.

In order to best focus a test on the particular function in question, testing mocks are often helpful replacements
for the functions' parameters. By creating Mocks that resemble those parameter types, but with quick and
lightweight implementations of their interface methods, we can guarantee and, in the case of spies, verify
specific interactions occur between the function and its parameters. If these events don't occur, then we can
fail the tests. In Go, the standard library also includes one such helper, `httptest`, which provides
convenient test helpers for faking external API calls that are not the focus of the test at hand.

#### Useful Testing Commands
- `go test` - The default command for running tests. It typically accepts any packages you've created as args
and also offers a `-v` flag for detailed logging
 - `go test ./...` is probably the best way to run tests since it'll recursively
 run any tests found in any subdirectories
 - `go list ./...` can similarly be used to list ALL packages found in subdirectories
 which makes it easier to run tests for specific packages in your project
 - `go test ./... -coverpkg=./...` can be used to get testing code coverage for your entire project


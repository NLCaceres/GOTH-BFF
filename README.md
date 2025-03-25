# GOTH-BFF - A Go Backend for a Templ-HTMX Frontend

An all around learning experience where I intend to explore Go in greater depths
as well as evaluate the many tools available in the Go ecosystem, all while I get
a better grasp on NeoVim as an IDE, personalizing it to my needs.

The goal of the project is to build a Go backend comparing the developer experience
of the Echo framework, the Chi Router and the Go Standard Library. This Go backend
will act as a gateway to a GraphQL API while also serving a UI based on Templ HTML
Components as the front-end to clients. The Templ UI will feature the HTMX library
as a means of triggering server-side re-renders as the user interacts with elements
on the page. Elements with HTMX attributes issue requests to the backend not only
to get new data, but to get a set of just the necessary components ready to swap
into place in the layout, already filled with all relevant data.

## Step by Step Plans & Layout of the Repo

- `Main` branch will always be the most up to date version of the project
- Each numbered branch represents the next major step in development
  - Branches with letters after the number represent small but important steps
  worth highlighting while developing a particular feature

0. The repo with a README, simple file structure, and basic dependencies
1. [Echo framework](https://github.com/labstack/echo) to set up the page routing
2. [Templ layout templates](https://github.com/a-h/templ) to set up basic HTML
structure of each page
3. [HTMX](https://htmx.org/docs/#introduction) to allow partial page updates in
relation to Nav bar or tab bar
4. Search bar to make POST requests to this backend to get more customized responses
from the underlying GraphQL API so results can be dynamically listed in the UI
5. Swap out Echo for a simpler [Chi implementation](https://github.com/go-chi/chi)
of the page routing
6. Swap out Chi for an even simpler [standard library implementation](https://pkg.go.dev/net/http#ServeMux)
   - [Recent improvements to Go's router](https://go.dev/blog/routing-enhancements)
   - [Advanced patterns in Go standard library routing video](https://www.youtube.com/watch?v=H7tbjKFSg58)
7. Dockerize for easy building, sharing and running

### Helpful Go Commands

- `make run` - Runs `main()` in `cmd/main.go` to setup server
  - `make watch` - Runs `main()` via `wgo` to restart server on code updates
  - All commands created in a `Makefile` can be run similarly by prepending
  `make` before the name of the command, e.g. `make lint`, `make build`, etc
- `go get <github-url>` - Install new dependency
  - `go get <github-url>@none` - Removes dependency
  - `go get .` - Finds external dependencies in current package files to install
  - `go get -tool <github-url>` - Installs tool dependencies
    - Adds `tool` directive to `go.mod` for pkgs like linters/formatters/builders
    - Pre-Go v1.24 tool dependencies could be split off into a `/internal/tools`
    package with a `tools.go` file importing the tools and a `go.mod` requiring them.
    BUT, this would require a `Makefile` command to globally install the runnables
- `go mod init <my-github/root-folder>`
  - Run in root folder to set up dependency tracking creating `go.mod`, i.e.
  `go mod init github.com/NLCaceres/goth-example` sets up a "goth-example" module
- `go mod tidy` - Clean up and optimize dependencies in `go.mod` and `go.sum` files
  - VERY useful since it can it act similarly to `go get .` and help install dependencies
  you may have missed, remove unused dependencies, downgrade or upgrade them, etc.
- `go run cmd/main.go` - Runs `main()` inside the cmd folder's `main.go` file
  - To run the module with live reloading, install `github.com/bokwoon95/wgo`
  and convert `go run` to `wgo run`
- `go install <github-url/external-module>@latest`
  - Installs a module's globally runnable executable
  - DOESN'T affect your local module as long as a version is included, e.g. "@latest"

### Testing

Golang testing is pretty awesome! It is particularly awesome, because everything
you need is baked into Go's standard library! Testing files end with `_test.go`,
and all test functions start with `Test` and include a parameter `t` of type `*testing.T`
(or sometimes of type `*testing.B` or `*testing.F` for Benchmarks or Fuzzing respectively).

For funcs with lots of cases to test, Golang uses table testing as a means of covering
all of these test cases from a single parent test func, structuring the test to be
very simple, readable and quick to write.

In order to best focus a test on the particular function in question, testing mocks
are often helpful replacements for the functions' parameters. By creating Mocks that
resemble those parameter types, but with quick and lightweight implementations of
their interface methods, we can guarantee and, in the case of spies, verify specific
interactions occur between the function and its parameters. If these events don't
occur, then we can fail the tests. In Go, the standard library also includes one
such helper, `httptest`, which provides convenient test helpers for faking external
API calls that are not the focus of the test at hand.

#### Useful Testing Commands

- `go test` - The default command for running tests. It typically accepts any
packages you've created as args and also offers a `-v` flag for detailed logging
  - `go test ./...` is probably the best way to run tests since it'll recursively
 run any tests found in any subdirectories
  - `go list ./...` can similarly be used to list ALL packages found in subdirectories
 which makes it easier to run tests for specific packages in your project
  - `go test ./... -coverpkg=./...` can be used to get testing code coverage for
  your entire project

# allez
Allez is the french word for go. This is where I learn Go, a
programming language.

## Programming Environment Setup

### Installation
Install Go for MacOS via the [package installer][installer]. The
package installs the Go distribution to `/usr/local/go`. The package
puts the `/usr/local/go/bin` directory in your `PATH` environment
variable.

### The `GOPATH` environment variable
The `GOPATH` is used to resolve import statements. The `GOPATH`
environment variable lists places (directory path) to look for Go
code. If not set, `GOPATH` defaults to `~/go`. To see the value of
`GOPATH`, run `go env GOPATH`. The Go path is often called the
`workspace` of Go. Each directory in `GOPATH` must have a prescribed
structure. [More on GOPATH][gopath]

## References
- [The Go Programming Language Specification][spec]
- [Xahlee's Golang Tutorial][xah]
- [A Tour of Go][tour]
- [The Go Programming Language (book)][gopl]
- [Go Concurrency Patterns (slides)][patterns]

[spec]: https://golang.google.cn/ref/spec
[xah]: http://xahlee.info/golang/golang_index.html
[tour]: https://tour.golang.org/list
[gopl]: https://www.gopl.io
[patterns]: https://talks.golang.org/2012/concurrency.slide
[installer]: https://www.golang.org/dl
[gopath]: https://golang.org/cmd/go/#hdr-GOPATH_environment_variable

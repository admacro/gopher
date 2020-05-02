# allez
Allez is the french word for go. This is where I learn Go, a
programming language.

### Programming Environment Setup

#### Installation
Install Go for MacOS via the [package installer][installer]. The
package installs the Go distribution to `/usr/local/go`. The package
puts the `/usr/local/go/bin` directory in your `PATH` environment
variable.

#### The `GOPATH` environment variable
The `GOPATH` is used to resolve import statements. The `GOPATH`
environment variable lists places (directory path) to look for Go
code. If not set, `GOPATH` defaults to `~/go`. To see the value of
`GOPATH`, run `go env GOPATH`. The Go path is often called the
`workspace` of Go. Each directory in `GOPATH` must have a prescribed
structure. [More on GOPATH][gopath]

### WIP
- [The Go Memory Model][memory]
- [Frequently Asked Questions (FAQ)][faq]
- [Diagnose Go Programes][diag]
- [The Go Wiki][wiki]

### Done
- [Xahlee's Golang Tutorial][xah]
- [A Tour of Go][tour]
- [Go Concurrency Patterns (slides)][patterns]
- [Writing web application in Go][gowiki]
- [Object-oriented programming without inheritance][oop]
- [The Go Programming Language Specification][spec]
- [Effective Go][effective]

### References
- [Go Documentation][doc]
- [The Go Blog][goblog]
- [Package Documentation][pkgdoc]
- [Command Documentation][cmddoc]
- [The Go Programming Language (book)][gopl]

[cmddoc]: https://golang.org/doc/cmd
[diag]: https://golang.org/doc/diagnostics.html
[doc]: https://golang.org/doc/
[effective]: https://golang.org/doc/effective_go.html
[faq]: https://golang.org/doc/faq
[goblog]: https://blog.golang.org/index
[gopath]: https://golang.org/cmd/go/#hdr-GOPATH_environment_variable
[gopl]: https://www.gopl.io
[gowiki]: https://github.com/admacro/gowiki
[installer]: https://www.golang.org/dl
[memory]: https://golang.org/ref/mem
[oop]: https://yourbasic.org/golang/inheritance-object-oriented/
[patterns]: https://talks.golang.org/2012/concurrency.slide
[pkgdoc]: https://golang.org/pkg/
[spec]: https://golang.google.cn/ref/spec
[tour]: https://tour.golang.org/list
[wiki]: https://github.com/golang/go/wiki
[xah]: http://xahlee.info/golang/golang_index.html

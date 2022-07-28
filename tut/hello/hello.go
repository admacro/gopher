// https://go.dev/doc/tutorial/getting-started
// https://pkg.go.dev/rsc.io/quote/v4
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}

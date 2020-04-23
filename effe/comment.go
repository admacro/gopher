// To see the generated doc, copy this file to `../../bonjour/comment`,
// then run `godoc -http=localhost:6060`. After that, go to
// http://localhost:6060/pkg/github.com/admacro/bonjour/comment/
// Comment till this line will not be included in the document.

/*
   The package doc goes here.
   Every package should have a package comment, a block comment preceding
   the package clause
   For multi-file packages, the package comment only needs to be present
   in one file, and any one will do.
   The package comment should introduce the package and provide information
   relevant to the package as a whole.
   see https://golang.org/doc/effective_go.html#commentary

   Indented text is diplayed in a fixed-width font.
   usage:
       import comment

       var s string = comment.ExportedFunc(comment.Pi)
       var expType comment.ExportedType{"field value"}
*/
package comment

import (
	"fmt"
)

// Inside a package, any comment immediately preceding a top-level declaration
// serves as a doc comment for that declaration.
//
// Doc comment for group declartion.
const (
	// pi is the ratio of a circle's circumference to its diameter.
	Pi = 3.1415926
	// e is the base of the natural logarithm
	E = 2.71828
)

// ExportedType is a simple demonstrative struct type.
// It's exported and can be accessed from outside this package.
// Every exported (capitalized) name in a program should have a doc comment.
type ExportedType struct {
	field string
}

// ExportedFunc tasks an arg of float and return a string contains the arg.
// If every doc comment begins with the name of the item it describes,
// you can use the doc subcommand of the go tool and run the output through
// grep.
//
// Doc comments work best as complete sentences, which allow a wide variety
// of automated presentations. The first sentence should be a one-sentence
// summary that starts with the name being declared.
func ExportedFunc(arg float64) string {
	return fmt.Sprintf("arg: %v", arg)
}

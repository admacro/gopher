// https://golang.org/doc/effective_go.html#interfaces
// https://golang.org/pkg/sort/
// https://golang.org/pkg/sort/#Interface
// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#IsSorted
// https://golang.org/pkg/sort/#Reverse
// https://golang.org/pkg/sort/#Slice
package sorter

import "sort"

type sorterKind int

const (
	Name sorterKind = iota
	Price
	Stock
)

type Product struct {
	Name  string
	Price float64
	Stock int
}

type byName []Product

func (p byName) Len() int           { return len(p) }
func (p byName) Less(i, j int) bool { return p[i].Stock < p[j].Stock }
func (p byName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type byPrice []Product

func (p byPrice) Len() int           { return len(p) }
func (p byPrice) Less(i, j int) bool { return p[i].Price < p[j].Price }
func (p byPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type byStock []Product

func (p byStock) Len() int           { return len(p) }
func (p byStock) Less(i, j int) bool { return p[i].Stock < p[j].Stock }
func (p byStock) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// https://golang.org/doc/effective_go.html#generality
// If a type exists only to implement an interface and will never have
// exported methods beyond that interface, there is no need to export the
// type itself.
// Exporting just the interface makes it clear the value has no interesting
// behavior beyond what is described in the interface. It also avoids the
// need to repeat the documentation on every instance of a common method.
//
// sort.Interface is implemented by three types in this package.
// These three types have no exported methods beyond sort.Interface.
// Intead of creating three separate functions with each returning a
// corresponding concrete type of sort.Interface, only one function
// is needed if it returns sort.Interface.
// sort.Interface is an exported type in package sort, the above applys
// to any interface in this package, in which case you only need to export
// that interface, instead of the three concrete types.
func BySorter(sk sorterKind, ps []Product) sort.Interface {
	switch sk {
	case Name:
		return byName(ps)
	case Price:
		return byPrice(ps)
	case Stock:
		return byStock(ps)
	default:
		return byName(ps)
	}
}

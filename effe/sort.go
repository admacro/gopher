// https://golang.org/doc/effective_go.html#interfaces
// https://golang.org/pkg/sort/
// https://golang.org/pkg/sort/#Interface
// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#IsSorted
// https://golang.org/pkg/sort/#Reverse
// https://golang.org/pkg/sort/#Slice
package main

import (
	"fmt"
	"sort"
)

type Product struct {
	name  string
	price float64
	stock int
}

type ByPrice []Product

func (p ByPrice) Len() int           { return len(p) }
func (p ByPrice) Less(i, j int) bool { return p[i].price < p[j].price }
func (p ByPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type ByStock []Product

func (p ByStock) Len() int           { return len(p) }
func (p ByStock) Less(i, j int) bool { return p[i].stock < p[j].stock }
func (p ByStock) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	ps := []Product{
		Product{"shoes", 69.9, 200},
		Product{"pants", 19.9, 100},
		Product{"t-shirt", 29.9, 50},
	}
	fmt.Println("\nsort by price")
	fmt.Printf("before sort, sort.IsSorted => %v\n", sort.IsSorted(ByPrice(ps)))
	sort.Sort(ByPrice(ps))
	print(ps)
	fmt.Printf("after sort, sort.IsSorted => %v\n", sort.IsSorted(ByPrice(ps)))

	fmt.Println("\nreverse sort by price with sort.Reverse")
	sort.Sort(sort.Reverse(ByPrice(ps)))
	print(ps)

	fmt.Println("\nsort by stock")
	sort.Sort(ByStock(ps))
	print(ps)

	// temporary sort
	// inline closure as a custom Less function for temporary sort
	// defined method Less is ignored
	fmt.Println("\nreverse sort by stock with custom Less function")
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].stock > ps[j].stock
	})
	print(ps)

	fmt.Println("\nsort by name with custom Less function")
	byName := func(i, j int) bool { return ps[i].name < ps[j].name } // closure
	sort.Slice(ps, byName)
	print(ps)
}

func print(ps []Product) {
	for i, p := range ps {
		fmt.Printf("#%d %+v\n", i+1, p)
	}
}

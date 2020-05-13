package sorter

import (
	"fmt"
	"sort"
	"testing"
)

func TestSorter(t *testing.T) {
	ps := []Product{
		{"shoes", 69.9, 200},
		{"pants", 19.9, 100},
		{"t-shirt", 29.9, 50},
	}
	fmt.Println("\nsort by price")
	bySorter := BySorter(Price, ps)
	fmt.Printf("before sort, sort.IsSorted => %v\n", sort.IsSorted(bySorter))
	sort.Sort(bySorter)
	print(ps)
	fmt.Printf("after sort, sort.IsSorted => %v\n", sort.IsSorted(bySorter))

	fmt.Println("\nreverse sort by price with sort.Reverse")
	sort.Sort(sort.Reverse(bySorter))
	print(ps)

	fmt.Println("\nsort by stock")
	bySorter = BySorter(Stock, ps)
	sort.Sort(bySorter)
	print(ps)

	// temporary sort
	// inline closure as a custom Less function for temporary sort
	// defined method Less is ignored
	fmt.Println("\nreverse sort by stock with custom Less function")
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Stock > ps[j].Stock
	})
	print(ps)

	fmt.Println("\nsort by name with custom Less function")
	byName := func(i, j int) bool { return ps[i].Name < ps[j].Name } // closure
	sort.Slice(ps, byName)
	print(ps)
}

func print(ps []Product) {
	for i, p := range ps {
		fmt.Printf("#%d %+v\n", i+1, p)
	}
}

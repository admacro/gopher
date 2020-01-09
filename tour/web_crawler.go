package main

import (
	"fmt"
	"sync"
)

var fetchedUrls = make(map[string]int)
var mux sync.Mutex
var wg sync.WaitGroup

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// Fetch URLs in parallel. <--done
	// Don't fetch the same URL twice. <--done
	defer wg.Done()

	if depth <= 0 {
		return
	} else {
		mux.Lock()
		_, urlExist := fetchedUrls[url]
		mux.Unlock()

		if urlExist {
			return
		}
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
		wg.Add(1)
	}
	return
}

func main() {
	go Crawl("https://golang.org/", 4, fetcher)
	wg.Add(1)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		mux.Lock()
		fetchedUrls[url] = 0
		mux.Unlock()
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/cmd/": &fakeResult{
		"Command",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/cmd/godoc/",
			"https://golang.org/cmd/gofmt/",
		},
	},
	"https://golang.org/cmd/godoc/": &fakeResult{
		"godoc command",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/cmd/gofmt/": &fakeResult{
		"gofmt command",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
		},
	},
}

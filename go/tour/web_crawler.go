package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("web_crawler")

	Crawl("https://golang.org/", 4, fetcher)
}

type Fetcher interface {
	// Fetch(url string, safeMap *SafeMap) (body string, urls []string, err error)

	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLS in Parallel
	// TODO Don't fetch the same URL twice

	// safeMap := &SafeMap{valueMap: make(map[string]string)}
	if depth <= 0 {
		return
	}
	// body, urls, err := fetcher.Fetch(url, safeMap)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

type fakeFetcher map[string]*fakeResult

// type fakeFetcher SafeMap

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	// if res, ok := f.valueMap[url]; ok {

	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	// valueMap: {
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
	// },
}

// Our SafeCounter has a mutex to keep the key from race conditions
type SafeMap struct {
	mutex    sync.Mutex
	valueMap map[string]fakeResult
}

func (c *SafeMap) Add(key string, val string) {
	c.mutex.Lock()
	fmt.Println(key, val)
	if c.valueMap == nil {
		c.valueMap = map[string]fakeResult{}
	}
	c.valueMap[key] = fakeResult{body: val}
	c.mutex.Unlock()
}

func (c *SafeMap) SValue(key string) string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.valueMap[key].body
}

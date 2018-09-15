package main

import (
	"flag"
	"fmt"
)

/* TODO

-> Handle cases such as https://abc.org/community/ and https://abc.org/community optimally
   I looked at strings.TrimSuffix() but it degrades performance - almost by 5 seconds with 20 
   worker threads!
-> Better UNIT test cases
-> Handle robots.txt
-> This web crawler only supports URLS (i.e. <a href></a>). This could easily extended to support
   assets such as css files, scripts, images, logos.
-> Pictorial respresentation of site map
-> As always, cleaner code, and optimal data structures, and better coding standards!
-> Make the hashmap thread safe.
*/

func main() {
	// command line arguments
	var (
		wc = flag.Int("c", 1, "Number of threads/workers to handle requests");
		url = flag.String("url", "https://www.amazon.co.uk/", "Origin/Starting point");
	)

	flag.Parse()

	defer func() {
		fmt.Println("**Errors encountered and recovered = ", recover())
	}()

	crawler := NewWebCrawler(*url, *wc)
	err := crawler.Crawl();

	if err != nil {
		panic("FATAL: Can't crawl")
	}
}

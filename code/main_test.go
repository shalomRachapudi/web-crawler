package main 

import (
	"fmt"
	"time"
	"testing"
	"reflect"
)

// unit test for crawl() in crawler.go
func TestCrawl(t *testing.T) {
	start := time.Now()
	crawler := NewWebCrawler("https://shalom7blog.wordpress.com/", 20)
	err := crawler.Crawl()
	if err != nil {
		t.Error("Error crawling")
	}
	fmt.Println("Time Elapsed crawling https://shalom7blog.wordpress.com/ with 20 worker threads is ", time.Since(start))
}

// a dummy HTML file
var (
	htmlDoc = `<html>
				 <body>
					<a href="http://shalomray7.com/home">Home</a>
					<a href="http://shalomray7.com/contact">Contact</a>
					<a href="https://monzo.org">Monzo</a>
					<a href="https://www.facebook.com">FB</a>
					<a href="https://monzo.org">Monzo</a>
				 </body>
				</html>`
)

// unit test for parseWebPage() in html_parser.go
func TestParseWebPage(t *testing.T) {
	webPage := NewWebPage(htmlDoc, "http://shalomray7.com/home")
	webPageInfo := webPage.ParseWebPage()

	expectedWebPageInfo := WebPageInfo{url: "http://shalomray7.com/home", linksWithin: map[string]bool{"http://shalomray7.com/home":true, "http://shalomray7.com/contact":true}}
	equal := reflect.DeepEqual(expectedWebPageInfo.linksWithin, webPageInfo.linksWithin)
	if !equal {
		t.Error("Expected internal URLs and actual internal URLs are not same")
		t.Error("Actual URLs", webPageInfo.linksWithin, "Expected: ", expectedWebPageInfo.linksWithin);
	}

	if expectedWebPageInfo.url != webPageInfo.url {
		t.Error("Expected URL and actual URLS are not same")
	}
}

// unit test for download() in downloader.go
func TestDownload(t *testing.T) {
	downloader := Downloader{}

	htmlScript, err := downloader.Download("https://monzo.com/") 
	if err != nil {
		t.Error("Couldn't download https://monzo.com");
	}

	if len(htmlScript) == 0 {
		t.Error("Response is of 0 bytes")
	}
}
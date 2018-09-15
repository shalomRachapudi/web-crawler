package main

import (
	"regexp"
	"strings"
	"golang.org/x/net/html"
)

// constants
const HREF string = "href"
const A_TAG string = "a"

var REGEX = regexp.MustCompile(`https?:\/\/([\w\d])+(\.\w+)*`) // to get base URL

type WebPage struct {
	htmlScript string
	baseURL string //consistent part of the URL
	url string // complete URL
}

func NewWebPage(doc string, url string) *WebPage {
	return &WebPage{doc, REGEX.FindString(url), url}
}

type WebPageInfo struct {
	url string

	// links within this page
	// using set to remove duplicates
	linksWithin map[string]bool
}

func NewWebPageInfo(url string) *WebPageInfo {
	var wpInfo WebPageInfo
	wpInfo.url = url
	wpInfo.linksWithin = make(map[string]bool)

	return &wpInfo
}

// extract information from the page
func (this WebPage) ParseWebPage() (wpInfo WebPageInfo) {

	htmlDoc, err := html.Parse(strings.NewReader(this.htmlScript));

	if err != nil {
		return WebPageInfo{}
	}

	wpInfo = *NewWebPageInfo(this.url);
	wpInfo.linksWithin = this.filterAndRemoveDuplicates(this.internalLinks(htmlDoc, HREF), this.baseURL)

	return wpInfo;
}

func (this WebPage) filterAndRemoveDuplicates(urls []string, baseURL string) (map[string]bool) {
	validUrl := ""
	result := make(map[string]bool)
	for _, url := range urls {
		if strings.HasPrefix(url, "/") {
			validUrl = baseURL + url
		} else if strings.HasPrefix(url, baseURL) {
			validUrl = url
		} else {
			continue
		}

		if !result[validUrl] {
			result[validUrl] = true
		}
	}
	return result
}

func (this WebPage) internalLinks(n *html.Node, key string) (tempUrls []string) {
	return this.extractInternalLinks(n, key, tempUrls)
}

func (this WebPage) extractInternalLinks(n *html.Node, key string, urls []string) ([]string) {

	if n.Type == html.ElementNode && n.Data == A_TAG {
		for _, attr := range n.Attr {
			if attr.Key == key {
				urls = append(urls, attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		urls = this.extractInternalLinks(c, key, urls);
	}

	return urls;
}
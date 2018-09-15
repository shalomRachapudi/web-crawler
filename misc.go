package main

import "fmt"

func print(doc WebPageInfo) {
	fmt.Println("URL:  ", doc.url)
	fmt.Println("Links: ")
	
	for key, _ := range doc.linksWithin {
		fmt.Println("    > ", key)
	}
}
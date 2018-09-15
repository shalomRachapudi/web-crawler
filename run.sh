#!/bin/sh

# uncomment the following line to run tests
# go test
go run main.go crawler.go html_parser.go misc.go downloader.go -c 20 -url "https://www.amazon.co.uk/" | tee output.out

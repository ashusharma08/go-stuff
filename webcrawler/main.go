package main

import "github.com/esoptra/go-prac/webcrawler/webcrawler"

func main() {
	wg := &webcrawler.WebCrawler{}
	wg.Start("https://www.amazon.in/")
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adonovan/gopl.io/ch5/links"
)

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

var tokens = make(chan struct{}, 20)

func crawLimit(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int
	n++
	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	depth := 0
	for ; n > 0; n-- {
		list := <-worklist
		if depth <= 1 {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						worklist <- crawLimit(link)
					}(link)
				}
			}

		}
		depth++
	}
}

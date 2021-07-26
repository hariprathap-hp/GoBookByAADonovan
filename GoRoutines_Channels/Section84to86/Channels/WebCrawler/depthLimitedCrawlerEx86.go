package main

import (
	"TheGoProgLangBook/Functions/Sections_5_4Onwards/links"
	"flag"
	"fmt"
	"log"
)

var depth = flag.Int("depth", 1, "Limit the depth of the crawl")

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	flag.Parse()
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	var n, d int                     // number of pending sends to worklist

	n++
	// record items per depth
	counter := make([]int, *depth+2)
	counter[d] = n

	// Add command-line arguments to worklist.
	// skip depth parameter
	go func() {
		worklist <- flag.Args()
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist

		// drain worklist then close unseenLinks
		if d > *depth {
			continue
		}
		for _, link := range list {
			if !seen[link] {
				//fmt.Println(link)
				n++ // counter++
				counter[d+1]++

				seen[link] = true
				unseenLinks <- link
			}
		}
		if counter[d]--; counter[d] == 0 {
			d++
		}
	}
	close(unseenLinks)
}

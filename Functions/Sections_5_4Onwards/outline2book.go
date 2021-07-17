package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

const (
	url = "https://www.poetryfoundation.org/poets/william-wordsworth"
)

var depth = 0

func main() {
	//first fetch the url using http.Get
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP GET failed for url %s with error %v\n", url, err)
		return
	}

	//now the parse the html response
	node, node_err := html.Parse(resp.Body)
	//close the response body
	resp.Body.Close()

	if node_err != nil {
		fmt.Printf("HTML parsing failed with error %v\n", node_err)
	}

	visit(node, firstChild, secondChild)
}

func visit(node *html.Node, fc, sc func(*html.Node)) {

	if fc != nil {
		fc(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visit(c, firstChild, secondChild)
	}

	if sc != nil {
		sc(node)
	}

}

func firstChild(fn *html.Node) {
	if fn.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", fn.Data)
		depth++
	}
}

func secondChild(sn *html.Node) {
	if sn.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", sn.Data)
		depth--
	}
}

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

var url = "https://espncricinfo.com"

func main() {
	node, err := pingURL(url)
	if err != nil {
		fmt.Println(err)
	}

	res := getElemsbyTag(node, "h1", "h2", "h3", "h4")
	for _, link := range res {
		fmt.Println(link)
	}
}

var res_node []*html.Node

func getElemsbyTag(doc *html.Node, strs ...string) []*html.Node {
	for _, str := range strs {
		if doc.Type == html.ElementNode && doc.Data == str && doc.FirstChild != nil {
			res_node = append(res_node, doc.FirstChild)
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		res_node = getElemsbyTag(c, strs...)
	}
	return res_node
}

/*func getElemsbyTag(doc *html.Node, name ...string) (out []*html.Node) {
	pre := func(doc *html.Node) {
		for _, n := range name {
			if doc.Type == html.ElementNode && doc.Data == n && doc.FirstChild != nil {
				out = append(out, doc.FirstChild)
			}
		}
	}
	forEachNode(doc, pre, nil)
	return
}

func forEachNode(d *html.Node, pre, post func(doc *html.Node)) {
	if pre != nil {
		pre(d)
	}

	for c := d.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, nil)
	}
}*/

func pingURL(url string) (*html.Node, error) {
	resp, r_err := http.Get(url)
	if r_err != nil {
		return nil, fmt.Errorf("HTTP Get failed for the url %s with an error %v", url, r_err)
	}

	node, n_err := html.Parse(resp.Body)
	if n_err != nil {
		return nil, fmt.Errorf("parsing of the response failed with an error %v", n_err)
	}
	resp.Body.Close()
	return node, nil
}

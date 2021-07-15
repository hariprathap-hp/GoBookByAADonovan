package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
type Node struct {
Type	NodeType
Data	string
Attr	[]Attribute
FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

	func Parse(r io.Reader) (*Node, error)
*/
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Finding Links failed %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(s []string, doc *html.Node) []string {
	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, v := range doc.Attr {
			if v.Key == "href" {
				s = append(s, v.Val)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		s = visit(s, c)
	}
	return s
}

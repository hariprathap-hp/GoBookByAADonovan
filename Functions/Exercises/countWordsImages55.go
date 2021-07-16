package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://espncricinfo.com"
	countWordsAndImages(url)
}

func countWordsAndImages(url string) error {
	//request for the url
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP Request Failed with error : ", err)
		return err
	}

	if resp.StatusCode != 200 {
		fmt.Println("Recevied Error in response ", resp.StatusCode)
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	//parse the response body
	node, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	images, words := actualCounting(node)
	fmt.Println(images, words)
	return nil
}

func actualCounting(node *html.Node) (images, words int) {
	if node.Type == html.ElementNode {
		if node.Data == "script" || node.Data == "style" {
			return
		} else if node.Data == "img" {
			images++
		}
	} else if node.Type == html.TextNode {
		text := node.Data
		reader := strings.NewReader(text)
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}

		/* another method to scan word by word
		   text := strings.TrimSpace(n.Data)
		   for _, line := range strings.Split(text, "\n") {
		       if line != "" {
		           words += len(strings.Split(line, " "))
		           //fmt.Printf("%s %q %d\n", line, strings.Split(line, " "), len(strings.Split(line, " ")))
		       }
		   }
		*/
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		w, i := actualCounting(c)
		words += w
		images += i
	}
	return
}

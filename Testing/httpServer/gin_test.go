package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGin(t *testing.T) {
	//get a gin engine
	r := gin.Default()
	//load HTML files
	r.LoadHTMLGlob("templates/*")
	r.GET("/", indexPage)

	//create a new request
	req, _ := http.NewRequest("GET", "/", nil)
	//create new response recorder
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if !f(rec) {
		t.Fail()
	}
}

func f(w *httptest.ResponseRecorder) bool {
	// Test that the http status code is 200
	statusOK := w.Code == http.StatusOK
	fmt.Println("Status ok is", statusOK)

	// Test that the page title is "Home Page"
	// You can carry out a lot more detailed tests using libraries that can
	// parse and process HTML pages
	p, err := ioutil.ReadAll(w.Body)
	pageOK := err == nil && strings.Index(string(p), "<title>Hello Test</title>") > 0

	return statusOK && pageOK
}

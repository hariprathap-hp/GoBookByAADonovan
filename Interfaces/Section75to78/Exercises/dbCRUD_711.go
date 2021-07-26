package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

var db = make(map[string]dollars)

func main() {
	db = database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(list))
	mux.Handle("/price", http.HandlerFunc(price))
	mux.Handle("/create", http.HandlerFunc(create))
	mux.Handle("/update", http.HandlerFunc(update))
	mux.Handle("/delete", http.HandlerFunc(del))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func list(w http.ResponseWriter, req *http.Request) {

	var itemTable = template.Must(template.New("Items").Parse(`
<h1>Items</h1>
<table>
    <tr>
        <th> Item </th>
        <th> Price </th>
    </tr>
    {{ range $k, $v := . }}
        <tr>
            <td>{{ $k }}</td>
            <td>{{ $v }}</td>
        </tr>
    {{end}}
</table>
`))
	itemTable.Execute(w, db)
}
func price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%f\n", price)
}

func create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	_, ok := db[item]
	if !ok {
		v, _ := strconv.ParseFloat(price, 32)
		db[item] = dollars(v)
	}

	for i, v := range db {
		fmt.Fprintf(w, "%s: %f\n", i, v)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	v, _ := strconv.ParseFloat(price, 32)
	_, ok := db[item]
	if !ok {
		db[item] = dollars(v)
	} else {
		db[item] = dollars(v)
	}
	for i, v := range db {
		fmt.Fprintf(w, "%s: %f\n", i, v)
	}
}

func del(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	_, ok := db[item]
	if ok {
		fmt.Fprintf(w, "Item %q is deleted successfully\n", item)
		delete(db, item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	for i, v := range db {
		fmt.Fprintf(w, "%s: %f\n", i, v)
	}
}

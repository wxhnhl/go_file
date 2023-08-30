package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{map[string]dollars{"shoes": 50, "socks": 5}, sync.Mutex{}}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/update", http.HandlerFunc(db.update))
	log.Fatal(http.ListenAndServe("localhost:80", mux))

}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	R map[string]dollars
	sync.Mutex
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.R {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db.R[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 64)
	if err != nil {
		fmt.Fprintf(w, "invalid price")
	}
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	if _, ok := db.R[item]; ok {
		db.R[item] = dollars(price)
		fmt.Fprintf(w, "update success")
	} else {
		fmt.Fprintf(w, "invalid item")
	}
}

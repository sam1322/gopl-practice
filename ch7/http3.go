package main

import (
	"fmt"
	"log"
	"net/http"
)

//func main() {
//	db := database{"shoes": 50, "socks": 5}
//	mux := http.NewServeMux()
//	//mux.Handle("/list", http.HandlerFunc(db.list))
//	//mux.Handle("/price", http.HandlerFunc(db.price)) // here handleFunc is adaptor and is used  for type conversion
//	mux.HandleFunc("/list", db.list)
//	mux.HandleFunc("/price", db.price)
//	log.Fatal(http.ListenAndServe(":8000", mux))
//}

func main() {
	db := database{"shoes": 50, "socks": 5}
	//mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list))
	//mux.Handle("/price", http.HandlerFunc(db.price)) // here handleFunc is adaptor and is used  for type conversion
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

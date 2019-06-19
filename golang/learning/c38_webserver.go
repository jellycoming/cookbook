package main

import (
	"fmt"
	"log"
	"net/http"
)

// curl -X POST -d'id=1' 'http://127.0.0.1:8000/what?name=foo'
func baseHandler(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		return
	}
	// map[id:[1] name:[foo]]
	fmt.Printf("form: %v\n", r.Form)
	// 1
	fmt.Printf("form arg: %v\n", r.Form.Get("id"))
	// /what?name=foo
	fmt.Printf("URL: %v\n", r.URL)
	// /what
	fmt.Printf("path: %v\n", r.URL.Path)
	// map[name:[foo]]
	fmt.Printf("query: %v\n", r.URL.Query())
	// foo
	fmt.Printf("query arg: %v\n", r.URL.Query().Get("name"))
	_, err := fmt.Fprintf(w, "Hello World! %s\n", r.URL.Path[1:])
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", baseHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

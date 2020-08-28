package main

import (
	"KV_store/store"
	"fmt"
	"log"
	"net/http"
)

var st *store.Data

//ServeHTTP handes requests to the server.
func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		store.Get(st, "hello")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))

	case "POST":
		store.Post(st, "hello", []byte("hello"))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))

	case "DELETE":
		store.Delete(st, "hello")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
func init() {
	st = &store.Data{
		Data: map[string][]byte{},
	}
}
func main() {

	http.HandleFunc("/", ServeHTTP)
	fmt.Println("Listening at localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

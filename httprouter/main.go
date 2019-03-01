package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if n, err := fmt.Fprint(w, "Welcome!\n"); err != nil {
		fmt.Println("the length of string ", n)
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if n, err := fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name")); err != nil {
		fmt.Println("the length of string ", n)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8081", router))
}

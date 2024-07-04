package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-book-store/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

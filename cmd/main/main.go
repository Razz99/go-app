package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/razz99/go-app/pkg/routes"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9000",r))
}
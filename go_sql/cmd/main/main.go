package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"github.com/krishnaGauss/go-proj/go_sql/pkg/routes"
)

func main(){
	r:=mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	// http.HandleFunc("/", r)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
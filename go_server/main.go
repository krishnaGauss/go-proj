package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	err:=r.ParseForm()
	if(err!=nil){
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //points to index.html file in directory

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hello" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "Method is not found", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "Hello World")
	})

	port := ":8080"
	println("Server running at http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

}

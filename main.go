package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your request is %s!", r.URL.Path[1:])
	fmt.Printf("Request is %s \n", r.URL.Path[1:]) // logger for console
}

func main() {
	fmt.Println("Server is listening for http://localhost:8080/api/v1/")
	http.HandleFunc("/api/v1/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

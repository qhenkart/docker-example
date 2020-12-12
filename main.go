package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thank you Docker")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("service started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

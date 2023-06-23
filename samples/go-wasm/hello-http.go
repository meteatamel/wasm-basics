package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method, r.URL.Path)
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server started. Listening on :8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenAndServe error:%s ", err.Error())
	}
}

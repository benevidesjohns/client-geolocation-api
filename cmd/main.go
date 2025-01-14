package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func main() {
	http.HandleFunc("/", handleTest)

	log.Println("server starting... port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

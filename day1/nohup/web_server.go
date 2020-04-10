package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hanler request")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

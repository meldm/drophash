package main

import (
    // "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "Hi there!")
    http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS(":8081", "config/cert.pem", "config/key.pem", nil)
}
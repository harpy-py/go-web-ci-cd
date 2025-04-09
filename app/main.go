package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go DevOps project!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "OK")
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/health", healthHandler)
    fmt.Println("Server is running on port 5000...")
    http.ListenAndServe(":5000", nil)
}
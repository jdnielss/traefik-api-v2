package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Handling request to Root")
        log.Printf("Handling request to %s", r.URL.Path)
        fmt.Fprintf(w, "Hello, World! Slurp")
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Handling request to Health")
        fmt.Fprintf(w, "OK, Health")
    })

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %s\n", err)
    }
}

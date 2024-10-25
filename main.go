package main

import (
    "fmt"
    "net/http"
    "log"
    "strings"
)

// Middleware to trim the prefix from the request path
func trimPrefixHandler(next http.HandlerFunc, prefix string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Strip the prefix from the request path
        path := strings.TrimPrefix(r.URL.Path, prefix)
        r.URL.Path = path

        // Call the next handler with the modified request
        next(w, r)
    }
}

func main() {
    // Define the prefix to be trimmed
    prefix := "/apps/traefik/v2"

    // Define the root handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Handling request to Root")
        log.Printf("Handling request to %s", r.URL.Path)
        fmt.Fprintf(w, "Hello, World! Slurp")
    })

    // Define the health handler
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Handling request to Health")
        fmt.Fprintf(w, "OK, Health")
    })

    // Wrap the root handler with the middleware to trim the prefix
    http.Handle("/", trimPrefixHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        http.ServeMux.DefaultServeMux.ServeHTTP(w, r)
    }), prefix))

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %s\n", err)
    }
}

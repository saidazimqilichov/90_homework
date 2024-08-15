package main

import (
    "fmt"
    "net/http"
)

// healthHandler - bu /health endpoint uchun handler.
func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Server is healthy!")
}

func main() {
    http.HandleFunc("/health", healthHandler) // /health endpoint uchun handler o'rnatish

    fmt.Println("Server running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}

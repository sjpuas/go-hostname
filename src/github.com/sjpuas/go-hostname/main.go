package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    name, _ := os.Hostname()
    fmt.Fprintf(w, "Hello World! Processing by container %s!", name)
}

func main() {
    http.HandleFunc("/", handler)
    log.Printf("WebServer running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

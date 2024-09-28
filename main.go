package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.HandleFunc("/datetime", func(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now().Format("2006-01-02 15:04:05")
        fmt.Fprintf(w, "%s", currentTime)
    })

    http.ListenAndServe(":8080", nil)
}
package main

import (
    "encoding/json"
    "net/http"
    "time"
)

type TimeResponse struct {
    CurrentTime string `json:"current_time"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := TimeResponse{CurrentTime: time.Now().Format(time.RFC1123)}
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/time", handler)
    http.ListenAndServe(":8080", nil)
}

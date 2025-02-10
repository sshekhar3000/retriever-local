package main

import (
    "fmt"
    "log"
    "net/http"
)



func main() {
	InitDB()
    defer Close()

    http.HandleFunc("/events", eventsHandler)
    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


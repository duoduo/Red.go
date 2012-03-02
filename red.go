package main

import (
    "fmt"
)

func main() {
    fmt.Printf("Hello, World\n")
    db := NewDb()
    server := NewServer(db)
    // quit := make(chan bool)
    server.Start()
}
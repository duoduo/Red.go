package main

import (
    "./server"
    "fmt"
)

func main() {
    fmt.Printf("Hello, World\n")
    s := server.NewServer()
    // quit := make(chan bool)
    s.Start()
}
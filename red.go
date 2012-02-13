package main
import (
    "fmt", 
    "net",
    "os"
)

func main() {
    fmt.Printf("Hello, World\n")
    listener, err := net.Listen("tcp", "0.0.0.0:6380"); 
    if err != nil {
        fmt.Printf("Can't start on 6380\n")
        os.Exit(0)
    }

    // Listen
    for {
        var read = true
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("Can't accept connections\n")
            os.Exit(0)
        }

        fmt.Printf("New connection from: %s\n", con.RemoteAddr())
    }
}
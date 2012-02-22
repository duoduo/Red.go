package request

import(
    "fmt"
    "strings"
)

type Request struct {
    Argc int
    Argv []string
}

func NewRequestFromQuery(query string) {
    //r := new(Request)
    
    // Redis Unified Request Protocol
    if query[0:0] != "*" {
        // Err
    }
    
    newlinepos := strings.Index(query, "\r")
    argc := query[1:newlinepos]
    fmt.Printf("NEW LINE POS: %s", newlinepos)
    fmt.Printf("ARGC: %s", argc)
}
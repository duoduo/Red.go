package red

import(
    "fmt"
    "strings"
)

const (
    CR_BYTE byte = byte("\r")
    LF_BYTE byte = byte("\n")
    COUNT_BYTE byte = byte("*")
)

type Request struct {
    Argc int
    Argv []string
}

func NewRequestFromConn(conn net.Conn) {
    //r := new(Request)
    //n, err := c.Conn.Read(buf[:])
    //fmt.Printf("N bytes: ", n)
    //fmt.Printf("Buf: %q", buf)
    //if err != nil {
    //    
    //}
    
    var buf [READ_BUF]byte
    n, err := c.Conn.Read(buf[:])

    // Redis Unified Request Protocol
    if query[0:0] != "*" {
        // Err
    }
    
    newlinepos := strings.Index(query, "\r")
    argc := query[1:newlinepos]
    fmt.Printf("NEW LINE POS: %s", newlinepos)
    fmt.Printf("ARGC: %s", argc)
}
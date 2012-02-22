package client

import(
    "fmt" 
    "net"
    "./request"
)

const (
    READ_BUF = (1024*16)
)

type Client struct {
    Conn net.Conn
}

func NewClient(conn net.Conn) *Client {
    //conn.SetReadBuffer(READ_BUF)
    c := Client{Conn: conn}
    return &c
}

func (c *Client) ReadQuery() {
    var buf [READ_BUF]byte
    n, err := c.Conn.Read(buf[:])
    fmt.Printf("N bytes: ", n)
    fmt.Printf("Buf: %q", buf)
    if err != nil {
        
    }
    
    query := string(buf[0:n])
    fmt.Printf("Query: %s", query)
    
    request.NewRequestFromQuery(query)
}
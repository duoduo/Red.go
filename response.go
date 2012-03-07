package main

import(
    "net"
    //"bytes"
    "fmt"
    "strconv"
)

type Response struct {
    Conn net.Conn
}

func NewResponse(conn net.Conn) *Response {
    return &Response{Conn: conn}
}

func (r *Response) Send(data []byte) {
    r.Conn.Write(data)
}

func (r *Response) Ok() {
    r.Send([]byte("+OK\r\n"))
}

func (r *Response) SendBulk(data []byte) {
    n := len(data)
    fmt.Printf("Len: ", []byte(strconv.Itoa(n)))
    d := []byte{'$'}
    //d = append(d, byte('$'))
    d = append(d, []byte(strconv.Itoa(n))...)
    d = append(d, []byte("\r\n")...)
    d = append(d, data...)
    d = append(d, []byte("\r\n")...)
    fmt.Printf("SendBulk: %s", d)
    r.Send(d)
}
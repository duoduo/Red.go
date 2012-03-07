package main

import(
    "net"
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
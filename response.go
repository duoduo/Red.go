package main

import(
    "net"
    //"bytes"
    "fmt"
    "strconv"
)


var shared = map[string] []byte {
    "pong": []byte("+PONG\r\n"),
}

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
    // Status Reply
    r.Send([]byte("+OK\r\n"))
}

func (r *Response) Pong() {
    r.Send(shared["pong"])
}

func (r *Response) Nil() {
    // Nil Reply
    r.Send([]byte("$-1\r\n"))
}

func (r *Response) Error(message string) {
    r.Send([]byte(fmt.Sprintf("-ERR %s \r\n", message)));
}

func (r *Response) SendBulk(data []byte) {
    n := len(data)
    fmt.Printf("Len: ", []byte(strconv.Itoa(n)))
    // Num of Args
    d := []byte{'$'}
    d = append(d, []byte(strconv.Itoa(n))...)
    d = append(d, []byte("\r\n")...)
    // Argument Data
    d = append(d, data...)
    d = append(d, []byte("\r\n")...)
    fmt.Printf("SendBulk: %s", d)
    r.Send(d)
}
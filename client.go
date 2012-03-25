package main

import (
	"fmt"
	"net"
	//"bytes"
	"bufio"
	"strconv"
	//"os"
)

const (
	CR_BYTE    byte = byte('\r')
	LF_BYTE    byte = byte('\n')
	COUNT_BYTE byte = byte('*')
	ARG_BYTE   byte = byte('$')
	READ_BUF        = (1024 * 16)
	MAX_ARGC        = (1024 * 1024)
)

// Request
type Request struct {
	Argc uint64
	Argv [][]byte
}


// Client
type Client struct {
	Conn     net.Conn
	Request  *Request
	Response *Response
	Command  CommandFunc
	Db       *Db
	Server   *Server
}

func NewClient(server *Server, db *Db, conn net.Conn) *Client {
	c := Client{
		Server:   server,
		Db:       db,
		Conn:     conn,
		Response: NewResponse(conn),
	}
	return &c
}

func (c *Client) ReadRequest(reader *bufio.Reader) bool {
	var lineBuf []byte
	var isPrefix bool
	var err error
	lineBuf, isPrefix, err = reader.ReadLine()
	if err != nil || isPrefix || lineBuf[0] != COUNT_BYTE {
		return false
	}

	var request *Request = new(Request)
	// Validate num of args.
	request.Argc, err = strconv.ParseUint(string(lineBuf[1:]), 10, 64)
	//fmt.Printf("\n\nARGC: ", request.Argc)
	if err != nil || request.Argc > MAX_ARGC {
		return false
	}

	// 2 lines per arg, 1st line is command, 2 line is value.
	request.Argv = make([][]byte, request.Argc*2)
	var line uint = 0
	for {
		if reader.Buffered() <= 0 {
			break
		}

		lineBuf, isPrefix, err = reader.ReadLine()
		if lineBuf[0] != ARG_BYTE {
			request.Argv[line] = lineBuf
			// New line if isPrefix == false
			if !isPrefix {
				line++
			}
		}
	}

	c.Request = request
	return true
}

func (c *Client) ProcessRequest() {
	reader := bufio.NewReaderSize(c.Conn, READ_BUF)
	for {
		// Read will block until something is ready to be ready.
		ok := c.ReadRequest(reader)
		if !ok {
			continue
		}

		c.Command = CommandFromRequest(c.Request)
		c.Command(c)

		// For now, force close conn when there is no more data.
		b, _ := reader.Peek(1)
		if len(b) == 0 {
			break
		}
 	}

 	fmt.Printf("Closing conn %s\n", c.Conn.RemoteAddr())
 	c.Conn.Close()
}

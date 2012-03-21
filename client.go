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

func NewRequestFromConn(conn net.Conn) {
	//n, err := c.Conn.Read(buf[:])
	//fmt.Printf("N bytes: ", n)
	//fmt.Printf("Buf: %q", buf)
	//if err != nil {
	//    
	//}
	var request *Request = new(Request)
	var lineBuf []byte
	var isPrefix bool // isPrefix specifies is line was fully read.
	var err error
	reader := bufio.NewReaderSize(conn, READ_BUF)

	lineBuf, isPrefix, err = reader.ReadLine()
	if err != nil || isPrefix || lineBuf[0] != COUNT_BYTE {
		// log.Exit(1)
	}
	fmt.Printf("ReadLine: ", lineBuf)
	fmt.Printf("ReadLine String: %s", lineBuf)
	/*for n, err := reader.Read(readBuf[:]); err == nil {
	    if n == 0 {
	        break
	    }
	}*/

	// Validate num of args.
	argc, err := strconv.ParseUint(string(lineBuf[1:]), 10, 64)
	fmt.Printf("\n\nARGC: ", argc)
	if err != nil || argc > MAX_ARGC {

	}

	// 2 lines per arg, 1st line is command, 2 line is value.
	request.Argv = make([][]byte, argc*2)
	var line uint = 0
	for {
		lineBuf, isPrefix, err = reader.ReadLine()
		if len(lineBuf) <= 0 {
			break
		}

		fmt.Printf("\n\nBufLength: %d Line %d: %s", len(lineBuf), line, lineBuf)

		request.Argv[line] = lineBuf
		// New line if isPrefix == false
		if !isPrefix {
			line++
		}
	}

	// Loop through num args
	//var fullBuf [READ_BUF]byte
	//_, _ = reader.Read(fullBuf[:])
	//fmt.Printf("ReadLine String FULL: %s", fullBuf)

	// Allocate storage
	//object := db.NewStringObject()

	// 

	// Redis Unified Request Protocol
	/*if query[0] != "*" {
	      // Err
	  }

	  newlinepos := strings.Index(query, "\r")
	  argc := query[1:newlinepos]
	  fmt.Printf("NEW LINE POS: %s", newlinepos)
	  fmt.Printf("ARGC: %s", argc)*/
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
	//conn.SetReadBuffer(READ_BUF)
	c := Client{
		Server:   server,
		Db:       db,
		Conn:     conn,
		Response: NewResponse(conn),
	}
	return &c
}

func (c *Client) ReadRequest() bool {
	//var buf [READ_BUF]byte
	//n, err := c.Conn.Read(buf[:])
	//fmt.Printf("N bytes: ", n)
	//fmt.Printf("Buf: %q", buf)
	//if err != nil {
	//    
	//}

	//query := string(buf[0:n])
	//fmt.Printf("Query: %s", query)
	var request *Request = new(Request)
	var lineBuf []byte
	var isPrefix bool // isPrefix specifies is line was fully read.
	var err error
	reader := bufio.NewReaderSize(c.Conn, READ_BUF)

	lineBuf, isPrefix, err = reader.ReadLine()
	if err != nil || isPrefix || lineBuf[0] != COUNT_BYTE {
		return false
	}
	//fmt.Printf("ReadLine: ", lineBuf)
	//fmt.Printf("ReadLine String: %s", lineBuf)
	/*for n, err := reader.Read(readBuf[:]); err == nil {
	    if n == 0 {
	        break
	    }
	}*/

	// Validate num of args.
	request.Argc, err = strconv.ParseUint(string(lineBuf[1:]), 10, 64)
	//fmt.Printf("\n\nARGC: ", request.Argc)
	if err != nil || request.Argc > MAX_ARGC {

	}

	// 2 lines per arg, 1st line is command, 2 line is value.
	request.Argv = make([][]byte, request.Argc*2)
	var line uint = 0
	for {
		if reader.Buffered() <= 0 {
			break
		}

		lineBuf, isPrefix, err = reader.ReadLine()
		//fmt.Printf("ERROR: ", err)
		//fmt.Printf("\n\nBufLength: %d Line %d: %s", len(lineBuf), line, lineBuf)
		if lineBuf[0] != ARG_BYTE {
			//fmt.Printf("Setting ARGV%s", line)
			request.Argv[line] = lineBuf
			// New line if isPrefix == false
			if !isPrefix {
				line++
			}
		}
	}

	//fmt.Printf("\n\nREQUEST DUMP: %+v", request)
	c.Request = request
	return true
}

func (c *Client) ProcessRequest(mainCh chan int) {
	for {
		// Read will block until something is ready to be ready.
		ok := c.ReadRequest()
		if !ok {
			continue
		}

		c.Command = CommandFromRequest(c.Request)
		
		// Alert that we are starting processing.
		<- mainCh
		c.Command(c)
		go func() {
		// Alert that we are done!
		// Next goroutine can take over.
			mainCh <- 1
		}()
	}
}

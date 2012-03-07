package main

import(
    "os"
    "fmt" 
    "net"
    //"bytes"
    "bufio"
    "strconv"
)

const (
    CR_BYTE byte = byte('\r')
    LF_BYTE byte = byte('\n')
    COUNT_BYTE byte = byte('*')
    READ_BUF = (1024 * 16)
    MAX_ARGC = (1024 * 1024)
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
    var err os.Error
    reader, err := bufio.NewReaderSize(conn, READ_BUF)
    if err != nil {

    }

    lineBuf, isPrefix, err = reader.ReadLine()
    if err != nil || isPrefix || lineBuf[0] != COUNT_BYTE {

    }
    fmt.Printf("ReadLine: ", lineBuf)
    fmt.Printf("ReadLine String: %s", lineBuf)
    /*for n, err := reader.Read(readBuf[:]); err == nil {
        if n == 0 {
            break
        }
    }*/

    // Validate num of args.
    argc, err := strconv.Atoui64(string(lineBuf[1:]))
    fmt.Printf("\n\nARGC: ", argc)
    if err != nil || argc > MAX_ARGC {

    }

    // 2 lines per arg, 1st line is command, 2 line is value.
    request.Argv = make([][]byte, argc * 2)
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
    Conn net.Conn
    Request *Request
    Response *Response
    Command Command
    Db *Db
}

func NewClient(db *Db, conn net.Conn) *Client {
    //conn.SetReadBuffer(READ_BUF)
    c := Client{
        Db: db,
        Conn: conn,
        Response: NewResponse(conn),
    }
    return &c
}

func (c *Client) ReadRequest() {
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
    var err os.Error
    reader, err := bufio.NewReaderSize(c.Conn, READ_BUF)
    if err != nil {

    }

    lineBuf, isPrefix, err = reader.ReadLine()
    if err != nil || isPrefix || lineBuf[0] != COUNT_BYTE {

    }
    fmt.Printf("ReadLine: ", lineBuf)
    fmt.Printf("ReadLine String: %s", lineBuf)
    /*for n, err := reader.Read(readBuf[:]); err == nil {
        if n == 0 {
            break
        }
    }*/

    // Validate num of args.
    request.Argc, err = strconv.Atoui64(string(lineBuf[1:]))
    fmt.Printf("\n\nARGC: ", request.Argc)
    if err != nil || request.Argc > MAX_ARGC {

    }

    // 2 lines per arg, 1st line is command, 2 line is value.
    request.Argv = make([][]byte, request.Argc * 2)
    var line uint = 0
    for {
        if reader.Buffered() <= 0 {
            break
        }

        lineBuf, isPrefix, err = reader.ReadLine() 
        fmt.Printf("ERROR: ", err)
        fmt.Printf("\n\nBufLength: %d Line %d: %s", len(lineBuf), line, lineBuf)

        request.Argv[line] = lineBuf        
        // New line if isPrefix == false
        if !isPrefix {
            line++
        }
    }

    fmt.Printf("\n\nREQUEST DUMP: %+v", request)
    c.Request = request
}

func (c *Client) ProcessRequest(mainCh chan int) {
    for {
        // Read will block until something is ready to be ready.
        c.ReadRequest()
        c.Command = CommandFromRequest(c.Request)
        // Alert that we are starting processing.
        <- mainCh
        c.Command.Process(c)
        go func() {
            // Alert that we are done!
            // Next goroutine can take over.
            mainCh <- 1
        }()
        //fmt.Printf("BEFORE RECEIVING")
        //fmt.Printf("RECEIVED IT!")
        // Take out 
    }
}
package main

import (
    "fmt" 
    "net"
    "os"
    "runtime/pprof"
)

type Server struct {
    Db *Db
}

func NewServer(db *Db) *Server {
    return &Server{Db: db}
}

func (s *Server) Start() {
    listener, err := net.Listen("tcp", "0.0.0.0:6380")
    if err != nil {
        fmt.Printf("Can't start on 6380\n")
        os.Exit(1)
    }

    // Main Channel, "single threaded".
    // Allows for multiple clients to be initialized.
    // Only allows one client to actually read/write to db.
    mainCh := make(chan int)
    go func() {
        // Tell the first client it can run.
        mainCh <- 1
    }()
    
    // Listen
    fmt.Printf("Listening on port 6380\n")
    for {
        // var read = true
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("Can't accept connections\n")
            os.Exit(1)
        }

        //fmt.Printf("New connection from: %s\n", conn.RemoteAddr())
        go s.handleConn(conn, mainCh)
    }
}

func (s *Server) Stop() {
    pprof.StopCPUProfile()
    os.Exit(0)
}

func (s *Server) handleConn(conn net.Conn, mainCh chan int) {    
    c := NewClient(s, s.Db, conn)
    c.ProcessRequest(mainCh)
}
package main

import (
    "fmt" 
    "net"
    "os"
    "runtime"
    "runtime/pprof"
)

type Server struct {
    Db *Db
}

func NewServer(db *Db) *Server {
    return &Server{Db: db}
}

func (s *Server) Start() {
    connCh := make(chan net.Conn, 1000)
    go s.handleConn(connCh)
    listener, err := net.Listen("tcp", "0.0.0.0:6380")
    if err != nil {
        fmt.Printf("Can't start on 6380\n")
        os.Exit(1)
    }
    
    // Listen
    fmt.Printf("Listening on port 6380\n")
    for {
        // var read = true
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("Can't accept connections\n")
            os.Exit(1)
        }

        connCh <- conn
    }
}

func (s *Server) Stop() {
    stats := new(runtime.MemStats)
    runtime.ReadMemStats(stats)
    fmt.Printf("Stats: (%+v", stats)
    pprof.StopCPUProfile()
    os.Exit(0)
}

func (s *Server) handleConn(connCh chan net.Conn) {
    for {
        conn := <-connCh
        c := NewClient(s, s.Db, conn)
        c.ProcessRequest()
    }
}
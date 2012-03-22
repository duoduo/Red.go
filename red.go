package main

import (
    "fmt"
    "flag"
    "os"
    "log"
    "runtime"
    "runtime/pprof"
    //"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
    runtime.GOMAXPROCS(1)
    fmt.Printf("Hello, World\n")
    flag.Parse()
    if *cpuprofile != "" {
        fmt.Printf("CPU Profiling Enabled")
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
    }

    db := NewDb()
    server := NewServer(db)
    // quit := make(chan bool)
    server.Start()
}
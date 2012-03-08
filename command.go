package main

//import object
import (
    "fmt"
    //"./client"
    "strings"
)

type BaseCommand struct {

}

type Command interface {
    Process(client *Client)
}

type StringCommand struct {
    
}

func (comm *StringCommand) Process(client *Client) {
    fmt.Printf("\n\nSTRINGTYPE: %s", strings.ToLower(string(client.Request.Argv[0])))
    // Determine string command type.
    switch strings.ToLower(string(client.Request.Argv[0])) {
        case "set":
            comm.Set(client)
        case "get":
            comm.Get(client)
    }
    // Execute command
    // data := client.Db.get("123")
    // n, err := client.Db.set("123")
    // Execute callback
    // callback()
    // client.Conn.Write(data[:])
}

func (comm *StringCommand) Set(client *Client) {
    fmt.Printf("\n\nSettings: %s %s", client.Request.Argv[1], client.Request.Argv[2])
    client.Db.Set(client.Request.Argv[1], client.Request.Argv[2])
    // Reply
    client.Response.Ok()
}

func (comm *StringCommand) Get(client *Client) {
    buf := client.Db.Get(client.Request.Argv[1])
    fmt.Printf("\n\nGET RESPONSE: %s", string(buf))
    client.Response.SendBulk(buf)
    //_, _ = client.Conn.Write([]byte("+OK\r\n"))
    //client.Conn.Close()
}

func CommandFromRequest(r *Request) Command {
    return new(StringCommand)
}

//func (c *StringCommand) Get(key string) 

/*func Factory(t string) (Command) {
    fmt.Printf(t)
    switch t {
        case "string":
            //c := new (StringCommand)
    }
    
    return &new(StringCommand)
}*/
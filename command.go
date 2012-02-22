package red

//import object
import (
    //"fmt"
)

type Command struct {
    Name string
}

type StringCommand struct {
    Command
}

func (c *StringCommand) Get(key string) 

/*func Factory(t string) (Command) {
    fmt.Printf(t)
    switch t {
        case "string":
            //c := new (StringCommand)
    }
    
    return &new(StringCommand)
}*/
include $(GOROOT)/src/Make.inc



TARG=red
GOFILES=\
    db.go\
    server.go\
    response.go\
    client.go\
    command.go\
    red.go\
        
include $(GOROOT)/src/Make.cmd


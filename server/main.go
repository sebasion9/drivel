package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)
const (
    HOST = "localhost"
    LISTEN_PORT = "7272"
    TYPE = "tcp"
)
// Listens for incoming client messages
// @TODO
// - handle errors
func listen_cmd() error {
    listen, err := net.Listen(TYPE, HOST+":"+LISTEN_PORT)
    if err != nil {
        return &TCPError{
            Err : err,
            Msg : fmt.Sprintf("failed to create listener, \nHOST: %s:%s", HOST,LISTEN_PORT),
        }
    }

    log_msg("created listener","success")

    for {
        conn, err := listen.Accept()
        if err != nil {
            return &TCPError {
                Err : err,
                Msg : fmt.Sprintf("failed accepting connection"),
            }
        }
        buffer := make([]byte, 256)
        bytes_read, err := conn.Read(buffer)
        if err != nil {
            return &TCPError {
                Err : err,
                Msg : fmt.Sprintf("failed to read from conn, \nbytes read:%d", bytes_read),
            }
        }
        // event = string(buffer)
        // handle incoming event here
    }
    return nil
}

func log_msg(msg string, _type string) {
    log.SetFlags(log.Ldate | log.Ltime)
    log_msg := fmt.Sprintf("[%s] %s\n", strings.ToUpper(_type), msg)
    log.Print(log_msg)
}

func log_err(err error) {
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log_msg := fmt.Sprintf("\n[ERR] %s\n", err)
    log.Print(log_msg)
}

func main() {
    log_msg("starting server", "info")
    err := listen_cmd()
    log_err(err)
}











package main

import (
    "fmt"
    "os"
    "net"
    "log"
)
const ADDR = "127.0.0.1"
const PORT = "7777"
func help() {
    fmt.Printf("available commands are: \n")
    fmt.Printf("\t[p]ause - pauses the player\n")
    fmt.Printf("\t[n]ext - plays next song\n")
    fmt.Printf("\t[b]ack - plays previous song\n")
}
func main() {
    args := os.Args[1:]
    if(len(args) != 1) {
        fmt.Printf("\tplease provide 1 arg\n")
        fmt.Printf("\tgoify_client <cmd>\n")
        return
    }
    arg := args[0]
    if(arg == "-h" || arg == "--help" || arg == "help") {
        help()
        return
    }
    conn, err := net.Dial("tcp", ADDR + ":" + PORT)
    if (err != nil) {
        log.Fatalf("failed connecting to goify server at '%s:%s'\n",ADDR,PORT)
    }
    var buffer []byte
    if(arg == "-p" || arg == "p" || arg == "pause") {
        buffer = []byte("pause\n")
    }
    if(arg == "-n" || arg == "n" || arg == "next") {
        buffer = []byte("next\n")
    }
    if(arg == "-b" || arg == "b" || arg == "back") {
        buffer = []byte("previous\n")
    }
    _, err = conn.Write(buffer)
    conn.Close()
    if (err != nil) {
        log.Fatalf("failed writing data: %s", string(buffer))
    }
}

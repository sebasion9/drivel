package main

import "fmt"

type TCPError struct {
    Err error
    Msg string
}
func (e *TCPError) Error() string {
    return fmt.Sprintf("TCP error\ndetails: %s\nmsg: %s\n", e.Err, e.Msg)
}

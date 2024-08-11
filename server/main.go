package main

import (
	"log"
	"math"
	"net"
	"os"
	"strings"
	"time"
)
const (
    HOST = "localhost"
    PORT = "7777"
    TYPE = "tcp"
)
func refresh_loop(token JSONToken) {
    log.Println("[100%] refresh token routine started successfully")
    log.Println("[INFO] waiting for incoming events")
    start := time.Now();
    for true {
        // token refreshing
        if (time.Since(start) > time.Duration(token.Expires_in * int(math.Pow10(9)))) {
            token, err := refresh(token.Refresh_Token);
            if (err != nil ){
                pretty_error("couldnt refresh auth token", err)
                return
            }
            os.Setenv("spot_refr", token.Refresh_Token)
            os.Setenv("spot_token", token.Access_token)
            start = time.Now()
        }
    }
}
func poll_events() {
    log.Println("[80%] server process started successfully")
    player := Player{Body: ""}
    listen, err := net.Listen(TYPE, HOST+":"+PORT)
    if (err != nil) {
        pretty_error("failed to open tcp port", err)
    }
    for true {
        conn, err := listen.Accept()
        if (err != nil) {
            log.Fatalf("[ERR] tcp port error: \n%s\n", err)
        }
        buffer := make([]byte, 32)
        _, err = conn.Read(buffer)
        if (err != nil) {
            log.Printf("[ERR] failed to read bytes from tcp port error: \n%s\n", err)
        }
        log.Println("[INFO] received an event")
        valid_event := false
        // based on buffer contents, run stuff here
        player.Token = os.Getenv("spot_token")
        cmd := string(buffer)
        switch {
        case strings.Contains(cmd, "pause"):
            _, err := player.pause()
            if (err != nil) {
                pretty_error("",err)
            } else {
                valid_event = true
            }
        case strings.Contains(cmd, "next"):
            _, err := player.skip(false)
            if (err != nil) {
                pretty_error("",err)
            } else {
                valid_event = true
            }
        case strings.Contains(cmd, "previous"):
            _,err := player.skip(true)
            if (err != nil) {
                pretty_error("",err)
            } else {
                valid_event = true
            }
        default:
            log.Printf("[FAIL] received event is in incorrect format: %s\n", cmd)
        }
        if(valid_event) {
            log.Println("[INFO] received event processed successfully")
        }
    }
}
func main() {
    log.Println("[0%] running goify server")
    refresh_token := os.Getenv("spot_refr")
    token, err := refresh(refresh_token);
    if (err != nil) {
        pretty_error("couldnt fetch auth token", err)
        return
    }
    log.Println("[40%] authentication token fetched successfully")
    os.Setenv("spot_token", token.Access_token)

    go refresh_loop(token)
    poll_events()
}


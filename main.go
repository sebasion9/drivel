package main

import (
    "fmt"
    "math"
    "time"
    "os"
)
func main() {
    refresh_token := os.Getenv("spot_refr")
    token, err := refresh(refresh_token);
    if (err != nil) {
        pretty_error("couldnt fetch auth token", err)
        return
    }
    body, err := pause(token)
    if (err != nil) {
        pretty_error("failed pause action", err)
    }
    fmt.Println(body)

    start := time.Now();
    for true {
        // token refreshing
        if (time.Since(start) > time.Duration(token.Expires_in * int(math.Pow10(9)))) {
            token, err = refresh(token.Refresh_Token);
            if (err != nil ){
                pretty_error("couldnt refresh auth token", err)
                return
            }
            os.Setenv("spot_refr", token.Refresh_Token)
            start = time.Now()
        }
        // event listening logic
    }
}

func pretty_error(message string, err error) {
    fmt.Println("<!ERR")
    fmt.Printf("\t")
    fmt.Println(message)
    fmt.Printf("\t")
    fmt.Println(err)
    fmt.Println("ERR!>")
}

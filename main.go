package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)
const base_url = "https://api.spotify.com/v1"
func main() {
    token, err := auth();
    if (err != nil) {
        pretty_error("couldnt fetch auth token", err)
        return
    }

    // test
    client := &http.Client{};
    req, err := http.NewRequest("GET", base_url + "/search", nil)
    req.Header.Set("Authorization", "Bearer " + token.Access_token);
    q := req.URL.Query()
    q.Add("limit","5");
    q.Add("type","artist");
    q.Add("q","franzl");
    req.URL.RawQuery = q.Encode();
    res, err := client.Do(req);
    // test

    if (err != nil) {
        pretty_error("test query failed", err);
    }
    body,err := io.ReadAll(res.Body);
    if (err != nil) {
        pretty_error("reading bytes failed", err);
    }
    fmt.Println(res.Status);
    fmt.Println(string(body));
    return

    start := time.Now();
    for true {
        // token refreshing
        if (time.Since(start) > time.Duration(token.Expires_in * int(math.Pow10(9)))) {
            token, err = auth();
            if (err != nil ){
                pretty_error("couldnt fetch auth token", err)
                return
            }
            start = time.Now()
        }
        // event listening logic
    }
    fmt.Println(token)
}

func pretty_error(message string, err error) {
    fmt.Println("<!ERR")
    fmt.Printf("\t")
    fmt.Println(message)
    fmt.Printf("\t")
    fmt.Println(err)
    fmt.Println("ERR!>")
}

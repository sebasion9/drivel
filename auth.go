package main

import (
	"encoding/base64"
    "encoding/json"
	"fmt"
	"io"
	"net/http"
    "net/url"
	"os"
	"strings"
)

const spotifyURL = "https://accounts.spotify.com/api/token" 
func get_env() (string, string) {
    id := os.Getenv("client_id");
    secret := os.Getenv("client_secret");
    return id, secret;
}
type JSONToken struct {
    Access_token string `json:"access_token"`
    Token_type string `json:"token_type"`
    Expires_in int `json:"expires_in"`
}
func auth() (JSONToken, error) {
    var token JSONToken
    client_id, client_secret := get_env();
    client := &http.Client{};
    params := url.Values{}; 
    params.Set("grant_type", "client_credentials");

    req, err := http.NewRequest("POST", spotifyURL, strings.NewReader(params.Encode()));
    if err != nil {
        fmt.Printf("error creating POST request: %s\n", err)
        return token, err
    }
    auth_header_val := "Basic " + base64.RawStdEncoding.EncodeToString([]byte(client_id + ":" + client_secret));
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Authorization", auth_header_val)
    res, err := client.Do(req);

    if err != nil {
        return token, err
    }
    body_as_bytes, err := io.ReadAll(res.Body)
    if err != nil {
        return token, err
    }
    err = json.Unmarshal(body_as_bytes,&token);
    if (err != nil) {
        return token, err
    }

    return token, nil
}


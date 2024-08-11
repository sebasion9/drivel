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

const auth_url = "https://accounts.spotify.com/api/token" 

func get_env() (string, string, string) {
    id := os.Getenv("spot_client_id");
    secret := os.Getenv("spot_client_secret");
    code := os.Getenv("spot_code");
    return id, secret, code;
}
type JSONToken struct {
    Access_token string `json:"access_token"`
    Token_type string `json:"token_type"`
    Expires_in int `json:"expires_in"`
    Refresh_Token string `json:"refesh_token"`
}
func auth() (JSONToken, error) {
    var token JSONToken
    client_id, client_secret, code := get_env();
    client := &http.Client{};
    params := url.Values{}; 
    params.Set("grant_type", "authorization_code");
    params.Set("code", code);
    params.Set("redirect_uri", "http://localhost:3003/callback");

    req, err := http.NewRequest("POST", auth_url, strings.NewReader(params.Encode()));
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
    fmt.Println(string(body_as_bytes))
    return token, nil
}
func refresh(refresh_token string) (JSONToken, error) {
    var token JSONToken;
    client_id, client_secret, _:= get_env();
    client := &http.Client{};
    params := url.Values{}; 
    params.Set("grant_type", "refresh_token");
    params.Set("refresh_token", refresh_token);
    params.Set("client_id", client_id)
    auth_header_val := "Basic " + base64.RawStdEncoding.EncodeToString([]byte(client_id + ":" + client_secret));
    req, err := http.NewRequest("POST", auth_url, strings.NewReader(params.Encode()));
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Authorization", auth_header_val)
    res,err := client.Do(req)
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


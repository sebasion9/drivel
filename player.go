package main

import (
    "net/http"
    "io"
)

const base_url = "https://api.spotify.com/v1"
// build betters errors?
func pause(token JSONToken) (string, error) {
    body := ""
    client := &http.Client{}
    target_url := base_url + "/me/player/pause"
    req, err := http.NewRequest("PUT", target_url, nil)
    if (err != nil) {
	return body, err
    }
    req.Header.Set("Authorization", "Bearer " + token.Access_token)
    res,err := client.Do(req)

    if (err != nil) {
	return body, err
    }
    body_bytes, err := io.ReadAll(res.Body)
    if (err != nil) {
	return body, err

    }
    body = string(body_bytes)
    return body, nil
}

// skip right/left
// pause/resume
// volume up/down
// play a song from [search, playlists]
// toggle repeat/random

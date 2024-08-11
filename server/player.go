package main
import (
	"io"
	"net/http"
)
type Player struct {
    Body string
    Client http.Client
    Token string
}
const BASE_URL = "https://api.spotify.com/v1"

func build_request_no_body(method string, url string, access_token string) (*http.Request, error) {
    req, err := http.NewRequest(method, url, nil)
    if(err != nil) {
	return req, &BuildRequestError{
	    Method: "PUT",
	    Err : err,
	}
    }
    req.Header.Set("Authorization", "Bearer " + access_token)
    return req, err
}

func (p *Player) pause() (string, error) {
    target_url := BASE_URL + "/me/player/pause"
    req,err := build_request_no_body("PUT", target_url, p.Token)
    if (err != nil) {
	return p.Body, err
    }
    res,err := p.Client.Do(req)
    if (err != nil) {
	return p.Body, &DoRequestError{
	    Err : err,
	}
    }
    body_bytes, err := io.ReadAll(res.Body)
    if (err != nil) {
	return p.Body, err
    }
    p.Body = string(body_bytes)
    if (!res_ok(res.StatusCode)) {
	return p.Body, &NotOkResponseError{
	    Body: p.Body,
	    Code: res.StatusCode,
	}
    }
    return p.Body, nil
}
func (p *Player) skip(previous bool) (string, error) {
    target_url := BASE_URL + "/me/player/next"
    if(previous) {
	target_url = BASE_URL + "/me/player/previous"
    }
    req, err := build_request_no_body("POST", target_url, p.Token)
    if(err != nil) {
	return p.Body, err
    }
    res, err := p.Client.Do(req)
    if(err != nil) {
	return p.Body, &DoRequestError{
	    Err : err,
	}
    }
    body_bytes, err := io.ReadAll(res.Body)
    if(err != nil) {
	return p.Body, err
    }
    p.Body = string(body_bytes)
    if (!res_ok(res.StatusCode)) {
	return p.Body, &NotOkResponseError{
	    Body: p.Body,
	    Code: res.StatusCode,
	}
    }
    return p.Body, nil
}

// resume
// volume up/down
// play a song from [search, playlists]
// toggle repeat/random

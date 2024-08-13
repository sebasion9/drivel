package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)
type StateProc struct {
    Progress_Ms float64
    Is_Playing bool
}

type Player struct {
    Body string
    Client http.Client
    Token string
    StateRaw map[string]interface{}
    State StateProc
}
const BASE_URL = "https://api.spotify.com/v1"

func build_request_no_body(method string, url string, access_token string) (*http.Request, error) {
    req, err := http.NewRequest(method, url, nil)
    if(err != nil) {
	return req, &BuildRequestError{
	    Method: method,
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
func (p *Player) state() (string, error) {
    target_url := BASE_URL + "/me/player"
    req, err := build_request_no_body("GET", target_url, p.Token)
    if (err != nil) {
	return p.Body, err
    }
    res, err := p.Client.Do(req)
    if (err != nil) {
	return p.Body, &DoRequestError{
	    Err: err,
	}
    }
    body_bytes, err := io.ReadAll(res.Body)
    if (err != nil) {
	return p.Body, err
    }
    p.Body = string(body_bytes)
    if(!res_ok(res.StatusCode)) {
	return p.Body, &NotOkResponseError{
	    Body: p.Body,
	    Code: res.StatusCode,
	}
    }
    err = json.Unmarshal([]byte(p.Body), &p.StateRaw)
    if err != nil {
	return p.Body, fmt.Errorf("[ERR] failed to parse state json body: \n%s\n", err)
    }
    var progress_ms float64
    var ok bool
    if progress_ms, ok = p.StateRaw["progress_ms"].(float64); !ok {
	log.Printf("[ERR] progress_ms is not desired type (float64)\n")
    } else {
	p.State.Progress_Ms = progress_ms
    }
    var is_playing bool
    if is_playing, ok = p.StateRaw["is_playing"].(bool); !ok {
	log.Printf("[ERR] is_playing is not desired type (bool)\n")
    } else {
	p.State.Is_Playing = is_playing 
    }

    return p.Body,nil
}
func (p *Player) resume() (string, error) {
    p.Body = ""
    if (p.State.Is_Playing) {
	return p.Body, nil
    }
    target_url := BASE_URL + "/me/player/play"
    post_body := map[string]int{"position_ms" : int(p.State.Progress_Ms)}
    json_body,err  := json.Marshal(post_body)
    if (err != nil) {
	return p.Body, fmt.Errorf("[ERR] failed to parse post json body: \n%s\n", err)
    }
    req, err := http.NewRequest("PUT", target_url, bytes.NewReader(json_body))
    if(err != nil) {
	return p.Body, &BuildRequestError{
	    Method: "PUT",
	    Err : err,
	}
    }
    req.Header.Set("Authorization", "Bearer " + p.Token)
    req.Header.Set("Content-Type", "application/json")
    res, err := p.Client.Do(req)
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
    if(!res_ok(res.StatusCode)) {
	return p.Body, &NotOkResponseError{
	    Body: p.Body,
	    Code: res.StatusCode,
	}
    }
    return p.Body, nil
}

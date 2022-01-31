package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

type FriendlyHttpResponse struct {
	rawRes     *http.Response
	BodyParsed interface{}
}
func HttpGet(url string)(friendlyResponse FriendlyHttpResponse, err error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	friendlyResponse.rawRes = resp
	// Check if body need parse json
	if resp.Header != nil && resp.Header.Get("Content-Type") == "application/json"{
		// JSON Body
		err := json.Unmarshal(body, &friendlyResponse.BodyParsed)
		if err != nil {
			return FriendlyHttpResponse{rawRes: resp}, err
		}
		return friendlyResponse, err
	}
	matchString, err := regexp.MatchString("^text/\\w+$", resp.Header.Get("Content-Type"))
	if err != nil {
		return FriendlyHttpResponse{rawRes: resp}, err
	}
	if matchString {
		// Text Body
		friendlyResponse.BodyParsed = string(body)
		return friendlyResponse, err
	}
	// Raw Body
	friendlyResponse.BodyParsed = body
	return
}

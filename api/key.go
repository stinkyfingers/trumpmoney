package api

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"syscall/js"
)

// GetAPIKey determines whether to fetch locally or from S3 and makes the request
func GetAPIKey() (string, error) {
	location := js.Global().Get("location").String()
	if strings.Contains(location, "localhost") {
		return apiKeyRequest(location + "/apikey")
	}
	return apiKeyRequest("https://fecapikey.s3-us-west-1.amazonaws.com/apikey")
}

func apiKeyRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	key := string(bytes.Trim(b, "\n\t "))
	if strings.Contains(key, "404") {
		return "", errors.New(key)
	}
	return key, nil
}

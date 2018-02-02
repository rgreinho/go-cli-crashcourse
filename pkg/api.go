package api

import (
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.ipify.org?format=json"

// This functions calls the ipify service.
func Call() (string, error) {
	res, err := http.Get(baseURL)
	if err != nil {
		return "", err
	}

	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

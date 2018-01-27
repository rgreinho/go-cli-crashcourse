package ipify

import (
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.ipify.org?format=json"

// Getter performs an HTTP GET request.
type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

// Client is a client for the 'ipify.org' api.
type Client struct {
	HTTPClient Getter
}

// Get performs an HTTP GET request against the 'ipify.org' api.
func (i *Client) Get() (*http.Response, error) {
	httpClient := i.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	res, err := httpClient.Get(baseURL)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (i *Client) String(res *http.Response) (string, error) {
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

// GetIP retrieves the public IP of the clien.
func (i *Client) GetIP() (string, error) {
	res, err := i.Get()
	if err != nil {
		return "", err
	}

	return i.String(res)
}

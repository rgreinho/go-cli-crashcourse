package ipify

import (
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.ipify.org?format=json"

type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

type IpifyClient struct {
	HTTPClient Getter
}

// This functions calls the ipify service.
func (i *IpifyClient) Get() (*http.Response, error) {
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

func (i *IpifyClient) String(res *http.Response) (string, error) {
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

func (i *IpifyClient) GetIP() (string, error) {
	res, err := i.Get()
	if err != nil {
		return "", err
	}

	return i.String(res)
}

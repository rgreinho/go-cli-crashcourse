package ipify

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ipifyJSONResponse = `{"ip":"123.456.789.123"}`

type ClientMock struct {
}

func (m *ClientMock) Get(url string) (*http.Response, error) {
	response := &http.Response{
		Header:     make(http.Header),
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "application/json")
	response.Body = ioutil.NopCloser(strings.NewReader(ipifyJSONResponse))

	return response, nil
}

func TestIpifyClientGet_00(t *testing.T) {
	mock := &ClientMock{}
	c := Client{HTTPClient: mock}

	res, err := c.GetIP()
	assert.Equal(t, err, nil)
	assert.Equal(t, ipifyJSONResponse, res)
}

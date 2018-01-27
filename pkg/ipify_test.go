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
	c := IpifyClient{HTTPClient: mock}
	res, _ := c.Get()

	actual := res.Body
	expected := ioutil.NopCloser(strings.NewReader(ipifyJSONResponse))
	assert.Equal(t, expected, actual, "")
}

func TestString_00(t *testing.T) {
	mock := &ClientMock{}
	res, _ := mock.Get("fake")
	i := IpifyClient{}

	actual, _ := i.String(res)
	expected := ipifyJSONResponse
	assert.Equal(t, expected, actual, "")
}

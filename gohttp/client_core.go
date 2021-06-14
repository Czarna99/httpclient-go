package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(httpmethod string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(httpmethod, url, nil)
	if err != nil {
		return nil, errors.New("unable to create a request")
	}
	return client.Do(request)
}

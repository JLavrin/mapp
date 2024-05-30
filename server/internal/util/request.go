package util

import (
	"io"
	"net/http"
)

type Req struct {
	Url    string
	Method string
	Body   io.Reader
}

func Request[T any](payload Req) T {
	var result T
	req, err := http.NewRequest("GET", payload.Url, payload.Body)

	if err != nil {
		return result
	}

	req.Method = payload.Method
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return result
	}

	defer res.Body.Close()

	return result
}
